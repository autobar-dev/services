use crate::controllers;
use crate::app_context::Context;

use serde::{
  Serialize,
  Deserialize,
};
use actix_web::{
  put,
  web,
  Responder,
  HttpResponse,
};

#[derive(Debug, Deserialize)]
pub struct SetRateBody {
  from: String,
  to: String,
  rate: f64,
}

#[derive(Debug, Serialize)]
struct SetRateResponse {
  status: String,
  error: Option<String>,
}

#[put("/set")]
pub async fn set_route(data: web::Data<Context>, body: web::Json<SetRateBody>) -> impl Responder {
  if body.from.len() != 3 || body.to.len() != 3 {
    return HttpResponse::BadRequest().json(
      SetRateResponse {
        status: "error".to_string(),
        error: Some("currencies should be three-letter, ISO 4217-compliant codes".to_string()),
      }
    );
  }

  let from = body.from.to_owned().to_uppercase();
  let to = body.to.to_owned().to_uppercase();
  let rate = body.rate;

  let result = controllers::set_rate_controller(data.get_ref().clone(), from, to, rate).await;

  if result.is_err() {
    return HttpResponse::InternalServerError().json(
      SetRateResponse {
        status: "error".to_string(),
        error: Some(result.unwrap_err().message),
      }
    );
  }

  HttpResponse::Ok().json(
    SetRateResponse {
      status: "ok".to_string(),
      error: None,
    }
  )
}