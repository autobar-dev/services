use crate::controllers;
use crate::types;
use crate::app_context::Context;

use actix_web::{
  get,
  web,
  Responder,
  HttpResponse,
};
use serde::{
  Serialize,
  Deserialize,
};

#[derive(Debug, Deserialize)]
pub struct CurrencyQuery {
  code: String,
}

#[derive(Debug, Serialize)]
struct CurrencyResponse {
  status: String,
  data: Option<types::Currency>,
  error: Option<String>,
}

#[get("/")]
pub async fn currency_route(data: web::Data<Context>, query: web::Query<CurrencyQuery>) -> impl Responder {
  if query.code.len() != 3 {
    return HttpResponse::BadRequest().json(
      CurrencyResponse {
        status: "error".to_string(),
        data: None,
        error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
      }
    );
  }

  let code = query.code.to_uppercase();
  let currency = controllers::get_currency_controller(data.as_ref().to_owned(), code).await;

  if currency.is_err() {
    return HttpResponse::InternalServerError().json(
      CurrencyResponse {
        status: "error".to_string(),
        data: None,
        error: Some(currency.unwrap_err().message),
      }
    );
  }

  let currency = currency.unwrap();

  HttpResponse::Ok().json(
    CurrencyResponse {
      status: "ok".to_string(),
      data: Some(currency),
      error: None,
    }
  )
}