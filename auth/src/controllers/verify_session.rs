use actix_web::http;
use chrono::Utc;

use crate::types;
use crate::models;

pub async fn verify_session_controller(
    context: types::AppContext,
    session_id: uuid::Uuid,
    user_agent: Option<String>,
) -> Result<i32, types::RestError> {
    let session = models::SessionModel::get_by_session_id(context, session_id).await;

    if session.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::NOT_FOUND,
            "session not found"
        ));
    }

    let session = session.unwrap();

    if Utc::now() > session.valid_until {
        return Err(types::RestError::new(
            http::StatusCode::BAD_REQUEST,
            "session out of date"
        ));
    }

    if user_agent.is_some() && session.user_agent.is_some() && user_agent.unwrap() != session.user_agent.unwrap() { // check if config.allow_only_same_user_agent is true
        return Err(types::RestError::new(
            http::StatusCode::BAD_REQUEST,
            "user agent different than one creating the session"
        ));
    }

    Ok(session.user_id)
}
