use crate::models;
use crate::app_context;
use crate::error::RestError;

pub async fn get_all_currencies_controller(context: app_context::Context) -> Result<Vec<models::CurrencyModel>, RestError> {
  let context = context.clone();
  let enabled_currencies = models::CurrencyModel::get_all(context).await;

  if enabled_currencies.is_err() {
    return Err(RestError::new("Error fetching currencies".to_string()));
  }

  let enabled_currencies = enabled_currencies.unwrap();
  
  Ok(enabled_currencies)
}