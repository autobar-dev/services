use crate::models;
use crate::types;
use crate::types::SessionInfo;

use actix_web::http;
use actix_web::http::StatusCode;

pub async fn get_sessions_for_user(
    context: types::AppContext,
    user_email: String,
) -> Result<Vec<SessionInfo>, types::RestError> {
    let user = models::UserModel::get_by_email(context.clone(), user_email).await;

    if user.is_err() {
        return Err(types::RestError::new(
            StatusCode::NOT_FOUND,
            "user not found",
        ));
    }

    let user = user.unwrap();

    let _delete_sessions =
        models::SessionModel::delete_all_expired_for_user(context.clone(), user.id).await;

    let sessions = models::SessionModel::all_for_user(context.clone(), user.id).await;

    if sessions.is_err() {
        return Err(types::RestError::new(
            http::StatusCode::INTERNAL_SERVER_ERROR,
            "unknown error occured",
        ));
    }

    let mut sessions = sessions.unwrap();
    sessions.sort_by(|a, b| b.last_used.cmp(&a.last_used));

    let sessions_infos: Vec<SessionInfo> =
        sessions.into_iter().map(types::SessionInfo::from).collect();

    Ok(sessions_infos)
}
