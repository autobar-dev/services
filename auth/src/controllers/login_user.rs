use crate::types;
use crate::models;

use actix_web::http;
use bcrypt;
use chrono::DateTime;
use chrono::Duration;
use chrono::Utc;

pub async fn login_user_controller(
    context: types::AppContext,
    email: String,
    password: String,
    remember_me: bool,
    user_agent: Option<String>, 
) -> Result<String, types::RestError> {
    let user_not_found_response: types::RestError = types::RestError::new(
        http::StatusCode::NOT_FOUND,
        "user with the specified credentials not found"
    );

    let context_user_clone = context.clone();
    let user = models::UserModel::get_by_email(context_user_clone, email).await;

    if user.is_err() {
        return Err(user_not_found_response);
    }

    let user = user.unwrap();

    let does_password_match = bcrypt::verify(password, user.password.as_str());

    if does_password_match.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "could not verify password"
        ));
    }

    let does_password_match = does_password_match.unwrap();

    if !does_password_match {
        return Err(user_not_found_response);
    }
    
    let time_valid: Duration;

    match remember_me {
        true => time_valid = Duration::days(30),
        false => time_valid = Duration::hours(6),
    }

    let context_session_clone = context.clone();

    let valid_until: DateTime<Utc> = Utc::now() + time_valid;
    let session_id = models::SessionModel::new(context_session_clone, user.id, user_agent, valid_until).await;

    if session_id.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "could not create a session"
        ));
    }

    let session_id = session_id.unwrap();

    Ok(session_id.hyphenated().to_string())
}
