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
  enabled: String,
}

#[derive(Debug, Serialize)]
struct SetEnabledCurrencyResponse {
  status: &'static str,
  data: Option<Vec<types::EnabledCurrency>>,
  error: Option<String>,
}

#[put("/set-enabled")]
pub async fn set_enabled_route(data: web::Data<Context>, body: web::Json<SetEnabledCurrencyBody>) -> impl Responder {
  if body.code.len() != 3 {
    return HttpResponse::BadRequest().json(
      SetEnabledCurrencyResponse {
        status: "error",
        data: None,
        error: Some("currency should be a three-letter, ISO 4217-compliant code".to_string()),
      }
    );
  }

  let code = body.code.to_owned().to_uppercase();
  let enabled: bool;

  if body.enabled.len() == 0 {
    return HttpResponse::BadRequest().json(
      SetEnabledCurrencyResponse {
        status: "error",
        data: None,
        error: Some("enabled not provided".to_string()),
      }
    );
  } else {
    if body.enabled.to_lowercase() == "false" || body.enabled == "0" {
      enabled = false;
    } else if body.enabled.to_lowercase() == "true" || body.enabled == "1" {
      enabled = true;
    } else {
      return HttpResponse::BadRequest().json(
        SetEnabledCurrencyResponse {
          status: "error",
          data: None,
          error: Some("enabled can be either 0/1 or true/false".to_string()),
        }
      );
    }
  }

  let result = controllers::set_currency_enabled_controller(data.get_ref().clone(), code, enabled).await;

  if result.is_err() {
    return HttpResponse::InternalServerError().json(
      SetEnabledCurrencyResponse {
        status: "error",
        data: None,
        error: Some(result.unwrap_err().message),
      }
    );
  }

  HttpResponse::Ok().json(
    SetEnabledCurrencyResponse {
      status: "ok",
      data: None,
      error: None,
    }
  )
}