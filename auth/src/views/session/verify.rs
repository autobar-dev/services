use crate::types;
use crate::controllers::verify_session_controller;

use actix_web::{
    http::header,
    http,
    web,
    get,
    Responder,
    HttpResponse, HttpRequest,
};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct VerifyQuery {
    session_id: String,
}

#[derive(Serialize, Debug)]
struct VerifyResponseData {
    email: String,
}

#[derive(Serialize, Debug)]
struct VerifyResponse {
    status: String,
    error: Option<String>,
    data: Option<VerifyResponseData>,
}

#[get("/verify")]
pub async fn verify_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    query: web::Query<VerifyQuery>
) -> impl Responder {
    let user_agent_header = req.headers().get(header::USER_AGENT);
    let mut user_agent: Option<String> = None; 

    if user_agent_header.is_some() {
       let user_agent_value = user_agent_header
           .unwrap()
           .to_str()
           .unwrap_or("");

        if user_agent_value != "" {
            user_agent = Some(user_agent_value.to_string());
        }
    }

    let provided_uuid = uuid::Uuid::parse_str(query.session_id.as_str());

    if provided_uuid.is_err() {
        return HttpResponse::BadRequest().json(
            VerifyResponse {
                status: "error".to_string(),
                error: Some("could not parse session id".to_string()),
                data: None,
            }
        );
    }

    let provided_uuid = provided_uuid.unwrap();

    let user_email = verify_session_controller(
        data.get_ref().clone(),
        provided_uuid,
        user_agent
    ).await;

    if user_email.is_err() {
        let user_email_error = user_email.unwrap_err();

        return match user_email_error.status {
            http::StatusCode::NOT_FOUND => HttpResponse::NotFound().json(
                VerifyResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some("session not found".to_string()),
                }
            ),
            http::StatusCode::BAD_REQUEST => HttpResponse::BadRequest().json(
                VerifyResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some("request incorrect".to_string()),
                }
            ),
            _ => HttpResponse::InternalServerError().json(
                VerifyResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some("unknown error".to_string()),
                }
            ),
        };
    }

    let user_email = user_email.unwrap();

    HttpResponse::Ok().json(
        VerifyResponse {
            status: "ok".to_string(),
            error: None,
            data: Some(VerifyResponseData {
                email: user_email,
            }),
        }
    )
}
