use crate::app_context::Context;
use crate::controllers;

use actix_web::{patch, web, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
pub struct UpdateRateBody {
    from: String,
    to: String,
}

#[derive(Debug, Serialize)]
struct UpdateRateResponse {
    status: String,
    error: Option<String>,
}

#[patch("/update")]
pub async fn update_route(
    data: web::Data<Context>,
    body: web::Json<UpdateRateBody>,
) -> impl Responder {
    if body.from.len() != 3 || body.to.len() != 3 {
        return HttpResponse::BadRequest().json(UpdateRateResponse {
            status: "error".to_string(),
            error: Some("currencies should be three-letter, ISO 4217-compliant codes".to_string()),
        });
    }

    let from = body.from.to_owned().to_uppercase();
    let to = body.to.to_owned().to_uppercase();

    let result = controllers::update_rate_controller(data.get_ref().clone(), from, to).await;

    if result.is_err() {
        return HttpResponse::InternalServerError().json(UpdateRateResponse {
            status: "error".to_string(),
            error: Some(result.unwrap_err().message),
        });
    }

    HttpResponse::Ok().json(UpdateRateResponse {
        status: "ok".to_string(),
        error: None,
    })
}

