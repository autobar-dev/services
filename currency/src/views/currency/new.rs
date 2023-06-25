use crate::app_context::Context;
use crate::controllers;

use actix_web::{post, web, HttpResponse, Responder};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
pub struct NewCurrencyBody {
    code: String,
    name: String,
    minor_unit_divisor: i32,
    symbol: Option<String>,
    enabled: Option<bool>,
}

#[derive(Debug, Serialize)]
struct NewCurrencyResponse {
    status: String,
    error: Option<String>,
}

#[post("/new")]
pub async fn new_route(
    data: web::Data<Context>,
    body: web::Json<NewCurrencyBody>,
) -> impl Responder {
    if body.code.len() != 3 {
        return HttpResponse::BadRequest().json(NewCurrencyResponse {
            status: "error".to_string(),
            error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
        });
    }

    if body.name.is_empty() {
        return HttpResponse::BadRequest().json(NewCurrencyResponse {
            status: "error".to_string(),
            error: Some("name cannot be empty".to_string()),
        });
    }

    if body.minor_unit_divisor <= 0 {
        return HttpResponse::BadRequest().json(NewCurrencyResponse {
            status: "error".to_string(),
            error: Some("minor_unit_divisor has to have a value higher than 0".to_string()),
        });
    }

    let code = body.code.to_owned().to_uppercase();
    let name = body.name.to_owned();
    let minor_unit_divisor = body.minor_unit_divisor;
    let symbol = body.symbol.clone();
    let enabled = body.enabled.unwrap_or(true);

    let result = controllers::new_currency_controller(
        data.get_ref().clone(),
        code,
        name,
        minor_unit_divisor,
        symbol,
        enabled,
    )
    .await;

    if result.is_err() {
        return HttpResponse::InternalServerError().json(NewCurrencyResponse {
            status: "error".to_string(),
            error: Some(result.unwrap_err().message),
        });
    }

    HttpResponse::Ok().json(NewCurrencyResponse {
        status: "ok".to_string(),
        error: None,
    })
}
