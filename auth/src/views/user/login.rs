use crate::types;

use actix_web::{
    http::header,
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
    let mut user_agent: &str = ""; 

    if user_agent_header.is_some() {
       user_agent = user_agent_header
           .unwrap()
           .to_str()
           .unwrap_or("");
    }

    log::info!("user-agent={}", user_agent);

    let email = body.email.to_lowercase();
    let password = body.password.clone();
    let remember_me = body.remember_me.unwrap_or(false);

    log::info!("email={}, password={}, remember_me={}", email, password, remember_me);

    HttpResponse::Ok().json(
        LoginUserResponse {
            status: "ok".to_string(),
            error: None,
            data: Some(LoginUserResponseData {
                session_id: "yooo".to_string(),
            })
        }
    )
}
