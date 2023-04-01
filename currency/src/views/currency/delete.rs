use crate::controllers;
use crate::app_context::Context;

use serde::{
  Serialize,
  Deserialize,
};
use actix_web::{
  delete,
  web,
  Responder,
  HttpResponse,
};

#[derive(Debug, Deserialize)]
pub struct DeleteCurrencyBody {
  code: String,
}

#[derive(Debug, Serialize)]
struct DeleteCurrencyResponse {
  status: String,
  error: Option<String>,
}

#[delete("/delete")]
pub async fn delete_route(data: web::Data<Context>, body: web::Json<DeleteCurrencyBody>) -> impl Responder {
  if body.code.len() != 3 {
    return HttpResponse::BadRequest().json(
      DeleteCurrencyResponse {
        status: "error".to_string(),
        error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
      }
    );
  }

  let code = body.code.to_owned().to_uppercase();

  let result = controllers::delete_currency_controller(data.get_ref().clone(), code).await;

  if result.is_err() {
    return HttpResponse::InternalServerError().json(
      DeleteCurrencyResponse {
        status: "error".to_string(),
        error: Some(result.unwrap_err().message),
      }
    );
  }

  HttpResponse::Ok().json(
    DeleteCurrencyResponse {
      status: "ok".to_string(),
      error: None,
    }
  )
}