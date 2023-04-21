use crate::controllers::remove_session_controller;
use crate::types;

use actix_web::{cookie::Cookie, delete, http, web, HttpRequest, HttpResponse, Responder};
use serde::{Deserialize, Serialize};
use time::Duration;

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
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    body: Option<web::Json<RemoveBody>>,
) -> impl Responder {
    let provided_uuid: Result<uuid::Uuid, uuid::Error>;

    let uuid_from_body: Option<String> = match body.is_some() {
        true => Some(body.unwrap().session_id.clone()),
        false => None,
    };

    let uuid_from_cookies = req.cookie("session_id");

    if uuid_from_body.is_some() {
        provided_uuid = uuid::Uuid::parse_str(uuid_from_body.unwrap().as_str());
    } else if uuid_from_cookies.is_some() {
        provided_uuid = uuid::Uuid::parse_str(
            uuid_from_cookies
                .unwrap_or(Cookie::new("session_id", ""))
                .value(),
        );
    } else {
        return HttpResponse::BadRequest().json(RemoveResponse {
            status: "error".to_string(),
            error: Some("session_id is missing from both body and cookies".to_string()),
        });
    }

    if provided_uuid.is_err() {
        return HttpResponse::BadRequest().json(RemoveResponse {
            status: "error".to_string(),
            error: Some("could not parse session id".to_string()),
        });
    }

    let provided_uuid = provided_uuid.unwrap();

    let removed_session = remove_session_controller(data.get_ref().clone(), provided_uuid).await;

    if removed_session.is_err() {
        let removed_session_err = removed_session.unwrap_err();

        return match removed_session_err.status {
            http::StatusCode::NOT_FOUND => HttpResponse::NotFound().json(RemoveResponse {
                status: "error".to_string(),
                error: Some(removed_session_err.message),
            }),
            http::StatusCode::BAD_REQUEST => HttpResponse::BadRequest().json(RemoveResponse {
                status: "error".to_string(),
                error: Some(removed_session_err.message),
            }),
            _ => HttpResponse::InternalServerError().json(RemoveResponse {
                status: "error".to_string(),
                error: Some("unknown error occured".to_string()),
            }),
        };
    }

    let removal_cookie = Cookie::build("session_id", "inactive")
        .max_age(Duration::ZERO)
        .path("/")
        .http_only(true)
        .finish();

    HttpResponse::Ok()
        .cookie(removal_cookie)
        .json(RemoveResponse {
            status: "ok".to_string(),
            error: None,
        })
}
