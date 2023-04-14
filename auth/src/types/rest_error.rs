use actix_web::http::StatusCode;
use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct RestError {
    pub message: String,

    #[serde(skip)]
    pub status: StatusCode, 
}

impl RestError {
    pub fn new(status_code: StatusCode, message: &str) -> RestError {
        RestError {
            message: message.to_string(),
            status: status_code,
        }
    } 
}
