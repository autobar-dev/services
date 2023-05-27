use crate::controllers::login_module_controller;
use crate::types;

use actix_web::{
    cookie::Cookie, http, http::header, post, web, HttpRequest, HttpResponse, Responder,
};
use serde::{Deserialize, Serialize};
use time::Duration;

#[derive(Deserialize, Debug)]
pub struct LoginModuleBody {
    serial_number: String,
    private_key: String,
    remember_me: Option<bool>,
}

#[derive(Serialize, Debug)]
struct LoginModuleResponseData {
    session_id: String,
}

#[derive(Serialize, Debug)]
struct LoginModuleResponse {
    status: String,
    error: Option<String>,
    data: Option<LoginModuleResponseData>,
}

#[post("/login")]
pub async fn login_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    body: web::Json<LoginModuleBody>,
) -> impl Responder {
    if body.serial_number.is_empty() || body.private_key.is_empty() {
        return HttpResponse::BadRequest().json(LoginModuleResponse {
            status: "error".to_string(),
            error: Some("both serial number and private key should be provided".to_string()),
            data: None,
        });
    }

    let user_agent_header = req.headers().get(header::USER_AGENT);
    let mut user_agent: Option<String> = None;

    if user_agent_header.is_some() {
        let user_agent_value = user_agent_header.unwrap().to_str().unwrap_or("");

        if !user_agent_value.is_empty() {
            user_agent = Some(user_agent_value.to_string());
        }
    }

    let serial_number = body.serial_number.clone();
    let private_key = body.private_key.clone();
    let remember_me = body.remember_me.unwrap_or(false);

    let session_id = login_module_controller(
        data.get_ref().clone(),
        serial_number,
        private_key,
        remember_me,
        user_agent,
    )
    .await;

    if session_id.is_err() {
        let session_id_error = session_id.unwrap_err();

        return match session_id_error.status {
            http::StatusCode::NOT_FOUND => HttpResponse::NotFound().json(LoginModuleResponse {
                status: "error".to_string(),
                data: None,
                error: Some("module not found".to_string()),
            }),
            _ => HttpResponse::InternalServerError().json(LoginModuleResponse {
                status: "error".to_string(),
                data: None,
                error: Some("unknown error".to_string()),
            }),
        };
    }

    let session_id = session_id.unwrap();

    let mut session_cookie_builder = Cookie::build("session_id", session_id.clone())
        .path("/")
        .http_only(true)
        .secure(data.config.set_secure_cookies);

    if remember_me {
        session_cookie_builder = session_cookie_builder
            .max_age(Duration::seconds(data.config.remember_me_duration_seconds));
    }

    let main_domain = data.config.main_domain.clone();

    if !main_domain.is_empty() {
        session_cookie_builder = session_cookie_builder.domain(main_domain);
    }

    let session_cookie = session_cookie_builder.finish();

    HttpResponse::Ok()
        .cookie(session_cookie)
        .json(LoginModuleResponse {
            status: "ok".to_string(),
            error: None,
            data: Some(LoginModuleResponseData {
                session_id: session_id.clone(),
            }),
        })
}
