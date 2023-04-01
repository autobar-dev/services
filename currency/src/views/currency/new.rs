use crate::controllers;
use crate::app_context::Context;

use serde::{
  Serialize,
  Deserialize,
};
use actix_web::{
  post,
  web,
  Responder,
  HttpResponse,
};

#[derive(Debug, Deserialize)]
pub struct NewCurrencyBody {
  code: String,
  name: String,
  enabled: Option<bool>,
}

#[derive(Debug, Serialize)]
struct NewCurrencyResponse {
  status: String,
  error: Option<String>,
}

#[post("/new")]
pub async fn new_route(data: web::Data<Context>, body: web::Json<NewCurrencyBody>) -> impl Responder {
  if body.code.len() != 3 {
    return HttpResponse::BadRequest().json(
      NewCurrencyResponse {
        status: "error".to_string(),
        error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
      }
    );
  }

  if body.name.len() == 0 {
    return HttpResponse::BadRequest().json(
      NewCurrencyResponse {
        status: "error".to_string(),
        error: Some("name cannot be empty".to_string()),
      }
    );
  }

  let code = body.code.to_owned().to_uppercase();
  let name = body.name.to_owned();
  let enabled = body.enabled.unwrap_or(true);

  let result = controllers::new_currency_controller(data.get_ref().clone(), code, name, enabled).await;

  if result.is_err() {
    return HttpResponse::InternalServerError().json(
      NewCurrencyResponse {
        status: "error".to_string(),
        error: Some(result.unwrap_err().message),
      }
    );
  }

  HttpResponse::Ok().json(
    NewCurrencyResponse {
      status: "ok".to_string(),
      error: None,
    }
  )
}