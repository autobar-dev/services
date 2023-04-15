use crate::types;
use crate::controllers::remove_session_controller;

use actix_web::{
    web,
    delete,
    Responder,
    HttpResponse,
    http,
};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct RemoveBody {
    session_id: String,
}

#[derive(Serialize, Debug)]
struct RemoveResponse {
    status: String,
    error: Option<String>,
}

#[delete("/remove")]
pub async fn remove_route(
    data: web::Data<types::AppContext>,
    body: web::Json<RemoveBody>
) -> impl Responder {
    let provided_uuid = uuid::Uuid::parse_str(body.session_id.as_str());

    if provided_uuid.is_err() {
        return HttpResponse::BadRequest().json(
            RemoveResponse {
                status: "error".to_string(),
                error: Some("could not parse session id".to_string()),
            }
        );
    }

    let provided_uuid = provided_uuid.unwrap();

    let removed_session = remove_session_controller(
        data.get_ref().clone(),
        provided_uuid,
    ).await;

    if removed_session.is_err() {
        let removed_session_err = removed_session.unwrap_err();

        return match removed_session_err.status {
            http::StatusCode::NOT_FOUND => HttpResponse::NotFound().json(
                RemoveResponse {
                    status: "error".to_string(),
                    error: Some(removed_session_err.message) 
                }
            ),
            http::StatusCode::BAD_REQUEST => HttpResponse::BadRequest().json(
                RemoveResponse {
                    status: "error".to_string(),
                    error: Some(removed_session_err.message) 
                }
            ),
            _ => HttpResponse::InternalServerError().json(
                RemoveResponse {
                    status: "error".to_string(),
                    error: Some("unknown error occured".to_string()),
                }
            ),
        }
    }

    HttpResponse::Ok().json(
        RemoveResponse {
            status: "ok".to_string(),
            error: None,
        }
    )
}
