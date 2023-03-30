use crate::controllers;
use crate::types;
use crate::app_context::Context;

use serde::Serialize;

use actix_web::{
  get,
  web,
  Responder,
  HttpResponse,
};

#[derive(Debug, Serialize)]
struct EnabledCurrenciesResponse {
  status: String,
  data: Option<Vec<types::EnabledCurrency>>,
  error: Option<String>,
}

#[get("/enabled")]
async fn enabled_route(data: web::Data<Context>) -> impl Responder {
  let enabled_currencies = controllers::get_enabled_currencies_controller(data.get_ref().clone()).await;

  if enabled_currencies.is_err() {
    return HttpResponse::InternalServerError().json(
      EnabledCurrenciesResponse {
        status: "error".to_string(),
        data: None,
        error: Some(enabled_currencies.unwrap_err().message),
      }
    );
  }

  let enabled_currencies = enabled_currencies.unwrap();

  HttpResponse::Ok().json(
    EnabledCurrenciesResponse {
      status: "ok".to_string(),
      data: Some(enabled_currencies),
      error: None,
    }
  )
}