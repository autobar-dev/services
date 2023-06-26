use crate::controllers;
use crate::types;
use crate::types::CreateModuleResponse;

use actix_web::{post, web, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct RegisterModuleBody {
    serial_number: String,
}

#[derive(Serialize, Debug)]
struct RegisterModuleResponse {
    status: String,
    error: Option<String>,
    data: Option<CreateModuleResponse>,
}

#[post("/register")]
pub async fn register_route(
    data: web::Data<types::AppContext>,
    body: web::Json<RegisterModuleBody>,
) -> impl Responder {
    if body.serial_number.is_empty() {
        return HttpResponse::BadRequest().json(RegisterModuleResponse {
            status: "error".to_string(),
            error: Some("serial_number should be provided".to_string()),
            data: None,
        });
    }

    let serial_number = body.serial_number.clone();

    let create_module_response =
        controllers::register_module_controller(data.get_ref().clone(), serial_number).await;

    if create_module_response.is_err() {
        return HttpResponse::BadRequest().json(RegisterModuleResponse {
            status: "error".to_string(),
            error: Some("failed to register module".to_string()),
            data: None,
        });
    }

    HttpResponse::Ok().json(RegisterModuleResponse {
        status: "ok".to_string(),
        error: None,
        data: Some(create_module_response.unwrap()),
    })
}
