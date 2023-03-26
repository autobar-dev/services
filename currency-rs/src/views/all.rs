use crate::controllers;
use crate::error::RestError;
use crate::app_context::Context;

use actix_web::{
  get,
  web,
  Responder,
  HttpResponse,
};

#[get("/all")]
pub async fn all_route(data: web::Data<Context>) -> impl Responder {
  let all_currencies = controllers::get_all_currencies_controller(data.as_ref().to_owned()).await;

  if all_currencies.is_err() {
    return HttpResponse::InternalServerError().json(RestError::new("Error fetching currencies".to_string()));
  }

  let all_currencies = all_currencies.unwrap();

  HttpResponse::Ok().json(all_currencies)
}