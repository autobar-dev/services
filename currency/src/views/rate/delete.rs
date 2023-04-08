use crate::app_context::Context;
use crate::controllers;

use actix_web::{delete, web, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
pub struct DeleteRateBody {
    from: String,
    to: String,
}

#[derive(Debug, Serialize)]
struct DeleteRateResponse {
    status: String,
    error: Option<String>,
}

#[delete("/delete")]
pub async fn delete_route(
    data: web::Data<Context>,
    body: web::Json<DeleteRateBody>,
) -> impl Responder {
    if body.from.len() != 3 || body.to.len() != 3 {
        return HttpResponse::BadRequest().json(DeleteRateResponse {
            status: "error".to_string(),
            error: Some("currencies should be three-letter, ISO 4217-compliant codes".to_string()),
        });
    }

    let from = body.from.to_owned().to_uppercase();
    let to = body.to.to_owned().to_uppercase();

    let result = controllers::delete_rate_controller(data.get_ref().clone(), from, to).await;

    if result.is_err() {
        return HttpResponse::InternalServerError().json(DeleteRateResponse {
            status: "error".to_string(),
            error: Some(result.unwrap_err().message),
        });
    }

    HttpResponse::Ok().json(DeleteRateResponse {
        status: "ok".to_string(),
        error: None,
    })
}
