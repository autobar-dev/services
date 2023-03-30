use crate::controllers;
use crate::types;
use crate::app_context::Context;

use actix_web::{
  get,
  web,
  Responder,
  HttpResponse,
};
use serde::Serialize;

#[derive(Debug, Serialize)]
struct AllCurrenciesResponse {
  status: String,
  data: Option<Vec<types::Currency>>,
  error: Option<String>,
}

#[get("/all")]
pub async fn all_route(data: web::Data<Context>) -> impl Responder {
  let all_currencies = controllers::get_all_currencies_controller(data.as_ref().to_owned()).await;

  if all_currencies.is_err() {
    return HttpResponse::InternalServerError().json(
      AllCurrenciesResponse {
        status: "error".to_string(),
        data: None,
        error: Some(all_currencies.unwrap_err().message),
      }
    );
  }

  let all_currencies = all_currencies.unwrap();

  HttpResponse::Ok().json(
    AllCurrenciesResponse {
      status: "ok".to_string(),
      data: Some(all_currencies),
      error: None,
    }
  )
}