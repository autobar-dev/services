use crate::models;
use crate::types;
use crate::types::SessionInfo;

use actix_web::http;
use actix_web::http::StatusCode;

pub async fn get_sessions_for_client(
    context: types::AppContext,
    client_type: types::ClientType,
    client_identifier: String,
) -> Result<Vec<SessionInfo>, types::RestError> {
    let client_not_found = match client_type {
        types::ClientType::User => {
            models::UserModel::get_by_email(context.clone(), client_identifier.clone())
                .await
                .is_err()
        }
        types::ClientType::Module => {
            models::ModuleModel::get_by_serial_number(context.clone(), client_identifier.clone())
                .await
                .is_err()
        }
    };

    if client_not_found {
        return Err(types::RestError::new(
            StatusCode::NOT_FOUND,
            "client not found",
        ));
    }

    let _delete_sessions = models::SessionModel::delete_all_expired_for_client(
        context.clone(),
        client_type,
        client_identifier.clone(),
    )
    .await;

    let sessions = models::SessionModel::all_for_client(
        context.clone(),
        client_type,
        client_identifier.clone(),
    )
    .await;

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
