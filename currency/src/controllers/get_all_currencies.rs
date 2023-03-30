use crate::models;
use crate::types;
use crate::app_context;
use crate::types::RestError;

pub async fn get_all_currencies_controller(context: app_context::Context) -> Result<Vec<types::Currency>, RestError> {
  let context = context.clone();
  let enabled_currencies = models::CurrencyModel::get_all(context).await;

  if enabled_currencies.is_err() {
    return Err(RestError::new("Error fetching currencies".to_string()));
  }

  let enabled_currencies = enabled_currencies.unwrap();
  
  let mut enabled_currencies_response: Vec<types::Currency> = Vec::new();

  for enabled_currency in enabled_currencies {
    enabled_currencies_response.push(types::Currency::from(enabled_currency));
  }

  Ok(enabled_currencies_response)
}