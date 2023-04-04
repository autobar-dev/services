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
pub struct RemoteRateQuery {
  from: String,
  to: String,
}

#[derive(Debug, Serialize)]
struct RemoteRateResponse {
  status: String,
  data: Option<types::RemoteRate>,
  error: Option<String>,
}

#[get("/remote")]
pub async fn remote_route(data: web::Data<Context>, query: web::Query<RemoteRateQuery>) -> impl Responder {
  if query.from.len() != 3 || query.to.len() != 3 {
    return HttpResponse::BadRequest().json(
      RemoteRateResponse {
        status: "error".to_string(),
        data: None,
        error: Some("both from and to parameters should be three-letter, ISO 4217-compliant codes".to_string()),
      }
    );
  }

  let from = query.from.to_uppercase();
  let to = query.to.to_uppercase();

  let rate = controllers::get_remote_rate_controller(data.get_ref().clone(), from, to).await;

  if rate.is_err() {
    let rate_error = rate.unwrap_err();

    return HttpResponse::BadRequest().json(
      RemoteRateResponse {
        status: "error".to_string(),
        data: None,
        error: Some(rate_error.message),
      }
    );
  }

  let rate = rate.unwrap();

  HttpResponse::Ok().json(
    RemoteRateResponse {
      status: "ok".to_string(),
      data: Some(rate),
      error: None,
    }
  )
}