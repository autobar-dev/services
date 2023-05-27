use crate::models;
use crate::types;

use actix_web::http;

use chrono::DateTime;
use chrono::Duration;
use chrono::Utc;

pub async fn login_module_controller(
    context: types::AppContext,
    serial_number: String,
    private_key: String,
    remember_me: bool,
    user_agent: Option<String>,
) -> Result<String, types::RestError> {
    let module_not_found_response: types::RestError = types::RestError::new(
        http::StatusCode::NOT_FOUND,
        "module with the specified credentials not found",
    );

    let module = models::ModuleModel::get_by_serial_number(context.clone(), serial_number).await;

    if module.is_err() {
        return Err(module_not_found_response);
    }

    let module = module.unwrap();

    let does_private_key_match = bcrypt::verify(private_key, module.private_key.as_str());

    if does_private_key_match.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "could not verify private key",
        ));
    }

    let does_private_key_match = does_private_key_match.unwrap();

    if !does_private_key_match {
        return Err(module_not_found_response);
    }

    let time_valid: Duration = match remember_me {
        true => Duration::seconds(context.config.remember_me_duration_seconds),
        false => Duration::seconds(context.config.default_session_duration_seconds),
    };

    let context_session_clone = context.clone();

    let valid_until: DateTime<Utc> = Utc::now() + time_valid;
    let session_id = models::SessionModel::create(
        context_session_clone,
        types::ClientType::Module,
        module.serial_number,
        user_agent,
        valid_until,
    )
    .await;

    if session_id.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "could not create a session",
        ));
    }

    let session_id = session_id.unwrap();

    Ok(session_id.hyphenated().to_string())
}
