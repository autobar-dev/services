use crate::types;
use crate::models;

use actix_web::http;
use uuid::Uuid;

pub async fn remove_session_controller(
    context: types::AppContext,
    session_uuid: Uuid, 
) -> Result<(), types::RestError> {
    let removed_session = models::SessionModel::delete(context, session_uuid).await;

    if removed_session.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "unknown error occured",
        ));
    }

    let removed_session = removed_session.unwrap(); // affected rows

    if removed_session == 0 {
        return Err(types::RestError::new(
            http::StatusCode::NOT_FOUND,
            "session not found"
        ));
    }

    Ok(())
}
