use crate::controllers;
use crate::types;

use actix_web::{
    cookie::Cookie, http, http::header, post, web, HttpRequest, HttpResponse, Responder,
};
use serde::{Deserialize, Serialize};
use time::Duration;

#[derive(Deserialize, Debug)]
pub struct RegisterUserBody {
    email: String,
    password: String,
    auto_login: Option<bool>,
    remember_me: Option<bool>,
}

#[derive(Serialize, Debug)]
struct RegisterUserResponseData {
    session_id: String,
}

#[derive(Serialize, Debug)]
struct RegisterUserResponse {
    status: String,
    error: Option<String>,
    data: Option<RegisterUserResponseData>,
}

#[post("/register")]
pub async fn register_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    body: web::Json<RegisterUserBody>,
) -> impl Responder {
    if body.email.is_empty() || body.password.is_empty() {
        return HttpResponse::BadRequest().json(RegisterUserResponse {
            status: "error".to_string(),
            error: Some("both email and password should be provided".to_string()),
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

    let email = body.email.to_lowercase();
    let password = body.password.clone();
    let remember_me = body.remember_me.unwrap_or(false);
    let auto_login = body.auto_login.unwrap_or(false);

    let session_id = controllers::register_user_controller(
        data.get_ref().clone(),
        email,
        password,
        remember_me,
        auto_login,
        user_agent,
    )
    .await;

    if session_id.is_err() {
        let session_id_error = session_id.unwrap_err();

        return match session_id_error.status {
            http::StatusCode::BAD_REQUEST => {
                HttpResponse::BadRequest().json(RegisterUserResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some(session_id_error.message),
                })
            }
            http::StatusCode::INTERNAL_SERVER_ERROR => {
                HttpResponse::InternalServerError().json(RegisterUserResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some(session_id_error.message),
                })
            }
            _ => HttpResponse::InternalServerError().json(RegisterUserResponse {
                status: "error".to_string(),
                data: None,
                error: Some("unknown error".to_string()),
            }),
        };
    }

    if !auto_login {
        return HttpResponse::Ok().json(RegisterUserResponse {
            status: "ok".to_string(),
            error: None,
            data: None,
        });
    }

    let session_id = session_id.unwrap().unwrap(); // at this point both
                                                   // result is Ok and option is Some

    let session_id_cookie_clone = session_id.clone();
    let mut session_cookie_builder = Cookie::build("session_id", session_id_cookie_clone)
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
        .json(RegisterUserResponse {
            status: "ok".to_string(),
            error: None,
            data: match auto_login {
                true => Some(RegisterUserResponseData {
                    session_id,
                }),
                false => None,
            },
        })
}
