use crate::models;
use crate::app_context;
use crate::error::RestError;

use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct EnabledCurrencyResponse {
  code: String,
  name: String,
}

impl EnabledCurrencyResponse {
  pub fn from(enabled_currency: models::CurrencyModel) -> EnabledCurrencyResponse {
    EnabledCurrencyResponse {
      code: enabled_currency.code.to_string(),
      name: enabled_currency.name.to_string(),
    }
  }
}

pub async fn get_enabled_currencies_controller(context: app_context::Context) -> Result<Vec<EnabledCurrencyResponse>, RestError> {
  let context = context.clone();
  let enabled_currencies = models::CurrencyModel::get_all_enabled(context).await;

  if enabled_currencies.is_err() {
    return Err(RestError::new("Error fetching currencies".to_string()));
  }

  let enabled_currencies = enabled_currencies.unwrap();

  let mut enabled_currencies_response: Vec<EnabledCurrencyResponse> = Vec::new();

  for enabled_currency in enabled_currencies {
    enabled_currencies_response.push(EnabledCurrencyResponse::from(enabled_currency));
  }

  Ok(enabled_currencies_response)
}