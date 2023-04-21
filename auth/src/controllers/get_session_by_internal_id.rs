use crate::models;
use crate::types;

use actix_web::http;

pub async fn get_session_by_internal_id_controller(
    context: types::AppContext,
    internal_id: i32,
) -> Result<types::Session, types::RestError> {
    let session = models::SessionModel::get_by_internal_id(context, internal_id).await;

    if session.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "unknown error occured",
        ));
    }

    let session = session.unwrap();

    Ok(types::Session::from(session))
}
