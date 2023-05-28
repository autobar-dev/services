use crate::controllers::get_session_by_internal_id_controller;

use crate::controllers::remove_session_controller;
use crate::controllers::verify_session_controller;
use crate::types;
use crate::types::consts::INTERNAL_HEADER_NAME;

use actix_web::cookie::Cookie;
use actix_web::http::header;
use actix_web::HttpRequest;
use actix_web::{delete, http, web, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct RemoveByInternalIdBody {
    internal_id: i32,
    session_id: Option<String>,
}

#[derive(Serialize, Debug)]
struct RemoveByInternalIdResponse {
    status: String,
    error: Option<String>,
}

#[delete("/remove-by-internal-id")]
pub async fn remove_by_internal_id_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    body: web::Json<RemoveByInternalIdBody>,
) -> impl Responder {
    let user_agent_header = req.headers().get(header::USER_AGENT);
    let mut user_agent: Option<String> = None;

    if user_agent_header.is_some() {
        let user_agent_value = user_agent_header.unwrap().to_str().unwrap_or("");

        if !user_agent_value.is_empty() {
            user_agent = Some(user_agent_value.to_string());
        }
    }

    let context = data.get_ref().clone();

    let provided_uuid: Result<uuid::Uuid, uuid::Error>;

    let session_from_body = body.session_id.clone();
    let session_from_cookie = req.cookie("session_id");

    if session_from_body.is_some() {
        provided_uuid = uuid::Uuid::parse_str(session_from_body.unwrap().as_str());
    } else if session_from_cookie.is_some() {
        provided_uuid = uuid::Uuid::parse_str(
            session_from_cookie
                .unwrap_or(Cookie::new("session_id", ""))
                .value(),
        );
    } else {
        return HttpResponse::BadRequest().json(RemoveByInternalIdResponse {
            status: "error".to_string(),
            error: Some("session_id is missing from both query and cookies".to_string()),
        });
    }

    if provided_uuid.is_err() {
        return HttpResponse::BadRequest().json(RemoveByInternalIdResponse {
            status: "error".to_string(),
            error: Some("could not parse session id".to_string()),
        });
    }

    let provided_uuid = provided_uuid.unwrap();

    let internal_header = req.headers().get(INTERNAL_HEADER_NAME);

    let verify_session_data =
        verify_session_controller(context.clone(), provided_uuid, internal_header, user_agent)
            .await;

    if verify_session_data.is_err() {
        return HttpResponse::BadRequest().json(RemoveByInternalIdResponse {
            status: "error".to_string(),
            error: Some("cannot verify session".to_string()),
        });
    }

    let verify_session_data = verify_session_data.unwrap();

    let session = get_session_by_internal_id_controller(context.clone(), body.internal_id).await;

    if session.is_err() {
        return HttpResponse::BadRequest().json(RemoveByInternalIdResponse {
            status: "error".to_string(),
            error: Some("session not found".to_string()),
        });
    }

    let session = session.unwrap();

    if session.client_identifier != verify_session_data.client_identifier {
        return HttpResponse::Unauthorized().json(RemoveByInternalIdResponse {
            status: "error".to_string(),
            error: Some("you are not allowed to remove somebody else's session".to_string()),
        });
    }

    let removed_session = remove_session_controller(data.get_ref().clone(), session.id).await;

    if removed_session.is_err() {
        let removed_session_err = removed_session.unwrap_err();

        return match removed_session_err.status {
            http::StatusCode::NOT_FOUND => {
                HttpResponse::NotFound().json(RemoveByInternalIdResponse {
                    status: "error".to_string(),
                    error: Some(removed_session_err.message),
                })
            }
            http::StatusCode::BAD_REQUEST => {
                HttpResponse::BadRequest().json(RemoveByInternalIdResponse {
                    status: "error".to_string(),
                    error: Some(removed_session_err.message),
                })
            }
            _ => HttpResponse::InternalServerError().json(RemoveByInternalIdResponse {
                status: "error".to_string(),
                error: Some("unknown error occured".to_string()),
            }),
        };
    }

    HttpResponse::Ok().json(RemoveByInternalIdResponse {
        status: "ok".to_string(),
        error: None,
    })
}
