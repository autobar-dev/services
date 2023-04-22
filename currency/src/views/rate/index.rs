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
pub struct RateQuery {
  from: String,
  to: String,
  no_fetch: Option<bool>,
}

#[derive(Debug, Serialize)]
struct RateResponse {
  status: String,
  data: Option<types::Rate>,
  error: Option<String>,
}

#[get("/")]
pub async fn index_route(data: web::Data<Context>, query: web::Query<RateQuery>) -> impl Responder {
  if query.from.len() != 3 || query.to.len() != 3 {
    return HttpResponse::BadRequest().json(
      RateResponse {
        status: "error".to_string(),
        data: None,
        error: Some("currencies should be three-letter, ISO 4217-compliant codes".to_string()),
      }
    );
  }

  let from = query.from.to_uppercase();
  let to = query.to.to_uppercase();
  let no_fetch = query.no_fetch.unwrap_or(false);

  let rate = controllers::get_rate_controller(data.as_ref().to_owned(), from, to, no_fetch).await;

  if rate.is_err() {
    return HttpResponse::InternalServerError().json(
      RateResponse {
        status: "error".to_string(),
        data: None,
        error: Some(rate.unwrap_err().message),
      }
    );
  }

  let rate = rate.unwrap();

  HttpResponse::Ok().json(
    RateResponse {
      status: "ok".to_string(),
      data: Some(rate),
      error: None,
    }
  )
}
