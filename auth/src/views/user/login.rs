use crate::{types, controllers::login_user_controller};

use actix_web::{
    http::header,
    http,
    web,
    post,
    Responder,
    HttpResponse, HttpRequest,
};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct LoginUserBody {
    email: String,
    password: String,
    remember_me: Option<bool>,
}

#[derive(Serialize, Debug)]
struct LoginUserResponseData {
    session_id: String,
}

#[derive(Serialize, Debug)]
struct LoginUserResponse {
    status: String,
    error: Option<String>,
    data: Option<LoginUserResponseData>,
}

#[post("/login")]
pub async fn login_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    body: web::Json<LoginUserBody>
) -> impl Responder {
    if body.email.len() == 0 || body.password.len() == 0 {
        return HttpResponse::BadRequest().json(
            LoginUserResponse {
                status: "error".to_string(),
                error: Some("both email and password should be provided".to_string()),
                data: None,
            }
        );
    }

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

    let email = body.email.to_lowercase();
    let password = body.password.clone();
    let remember_me = body.remember_me.unwrap_or(false);

    let session_id = login_user_controller(
        data.get_ref().clone(),
        email,
        password,
        remember_me,
        user_agent
    ).await;

    if session_id.is_err() {
        let session_id_error = session_id.unwrap_err();

        return match session_id_error.status {
            http::StatusCode::NOT_FOUND => HttpResponse::NotFound().json(
                LoginUserResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some("user not found".to_string()),
                }
            ),
            _ => HttpResponse::InternalServerError().json(
                LoginUserResponse {
                    status: "error".to_string(),
                    data: None,
                    error: Some("unknown error".to_string()),
                }
            ),
        };
    }

    let session_id = session_id.unwrap();

    HttpResponse::Ok().json(
        LoginUserResponse {
            status: "ok".to_string(),
            error: None,
            data: Some(LoginUserResponseData {
                session_id,
            })
        }
    )
}
