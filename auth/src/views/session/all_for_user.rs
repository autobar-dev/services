use crate::{
    controllers::{get_sessions_for_user, verify_session_controller},
    types,
};

use actix_web::{cookie::Cookie, get, http::header, web, HttpRequest, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct AllForUserQuery {
    session_id: Option<String>,
}

#[derive(Serialize, Debug)]
pub struct AllForUserResponse {
    status: String,
    error: Option<String>,
    data: Option<Vec<types::SessionInfo>>,
}

#[get("/all-for-user")]
pub async fn all_for_user_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    query: web::Query<AllForUserQuery>,
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

    let session_from_query = query.session_id.clone();
    let session_from_cookie = req.cookie("session_id");

    if session_from_query.is_some() {
        provided_uuid = uuid::Uuid::parse_str(session_from_query.unwrap().as_str());
    } else if session_from_cookie.is_some() {
        provided_uuid = uuid::Uuid::parse_str(
            session_from_cookie
                .unwrap_or(Cookie::new("session_id", ""))
                .value(),
        );
    } else {
        return HttpResponse::BadRequest().json(AllForUserResponse {
            status: "error".to_string(),
            error: Some("session_id is missing from both query and cookies".to_string()),
            data: None,
        });
    }

    if provided_uuid.is_err() {
        return HttpResponse::BadRequest().json(AllForUserResponse {
            status: "error".to_string(),
            error: Some("could not parse session id".to_string()),
            data: None,
        });
    }

    let provided_uuid = provided_uuid.unwrap();

    let user_email = verify_session_controller(context.clone(), provided_uuid, user_agent).await;

    if user_email.is_err() {
        return HttpResponse::BadRequest().json(AllForUserResponse {
            status: "error".to_string(),
            error: Some("invalid session".to_string()),
            data: None,
        });
    }

    let user_email = user_email.unwrap();

    let session_infos = get_sessions_for_user(context.clone(), user_email).await;

    if session_infos.is_err() {
        return HttpResponse::InternalServerError().json(AllForUserResponse {
            status: "error".to_string(),
            error: Some("unknown error occured".to_string()),
            data: None,
        });
    }

    let session_infos = session_infos.unwrap();

    HttpResponse::Ok().json(AllForUserResponse {
        status: "ok".to_string(),
        error: None,
        data: Some(session_infos),
    })
}
