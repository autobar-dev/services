use actix_web::http;
use actix_web::http::header::HeaderValue;
use chrono::Utc;

use crate::models;
use crate::types;

#[derive(Clone, Debug)]
pub struct VerifySessionData {
    pub client_type: types::ClientType,
    pub client_identifier: String,
}

pub async fn verify_session_controller(
    context: types::AppContext,
    session_id: uuid::Uuid,
    internal_header: Option<&HeaderValue>,
    user_agent: Option<String>,
) -> Result<VerifySessionData, types::RestError> {
    let context_session_clone = context.clone();

    let session = models::SessionModel::get_by_session_id(context_session_clone, session_id).await;

    if session.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::NOT_FOUND,
            "session not found",
        ));
    }

    let session = session.unwrap();

    if Utc::now() > session.valid_until {
        let context_delete_clone = context.clone();
        let _delete_old_session =
            models::SessionModel::delete(context_delete_clone, session_id).await;

        return Err(types::RestError::new(
            http::StatusCode::BAD_REQUEST,
            "session out of date",
        ));
    }

    let context_config_clone = context.clone();

    if internal_header.is_none()
        && context_config_clone.config.allow_only_same_user_agent
        && user_agent.is_some()
        && session.user_agent.is_some()
        && user_agent.unwrap() != session.user_agent.unwrap()
    {
        return Err(types::RestError::new(
            http::StatusCode::BAD_REQUEST,
            "user agent different than one creating the session",
        ));
    }

    let context_update_last_used_clone = context.clone();

    let updated_rows =
        models::SessionModel::update_last_used(context_update_last_used_clone, session_id).await;

    if updated_rows.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "could not update last used",
        ));
    }

    Ok(VerifySessionData {
        client_type: session.client_type,
        client_identifier: session.client_identifier,
    })
}
