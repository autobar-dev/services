use crate::controllers;
use crate::error::RestError;
use crate::app_context::Context;

use actix_web::{
  get,
  web,
  Responder,
  HttpResponse,
};

#[get("/enabled")]
async fn enabled_route(data: web::Data<Context>) -> impl Responder {
  let enabled_currencies = controllers::get_enabled_currencies_controller(data.get_ref().clone()).await;

  if enabled_currencies.is_err() {
    return HttpResponse::InternalServerError().json(RestError::new("Error fetching currencies".to_string()));
  }

  let enabled_currencies = enabled_currencies.unwrap();

  HttpResponse::Ok().json(enabled_currencies)
}