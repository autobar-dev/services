use crate::controllers;
use crate::types;
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
pub struct SetEnabledCurrencyBody {
  code: String,
  enabled: bool,
}

#[derive(Debug, Serialize)]
struct SetEnabledCurrencyResponse {
  status: String,
  data: Option<Vec<types::EnabledCurrency>>,
  error: Option<String>,
}

#[put("/set-enabled")]
pub async fn set_enabled_route(data: web::Data<Context>, body: web::Json<SetEnabledCurrencyBody>) -> impl Responder {
  if body.code.len() != 3 {
    return HttpResponse::BadRequest().json(
      SetEnabledCurrencyResponse {
        status: "error".to_string(),
        data: None,
        error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
      }
    );
  }

  let code = body.code.to_owned().to_uppercase();
  let enabled = body.enabled;

  let result = controllers::set_currency_enabled_controller(data.get_ref().clone(), code, enabled).await;

  if result.is_err() {
    return HttpResponse::InternalServerError().json(
      SetEnabledCurrencyResponse {
        status: "error".to_string(),
        data: None,
        error: Some(result.unwrap_err().message),
      }
    );
  }

  HttpResponse::Ok().json(
    SetEnabledCurrencyResponse {
      status: "ok".to_string(),
      data: None,
      error: None,
    }
  )
}