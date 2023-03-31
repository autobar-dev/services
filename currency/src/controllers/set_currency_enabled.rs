use crate::models;
use crate::app_context;
use crate::types::RestError;

pub async fn set_currency_enabled_controller(context: app_context::Context, code: String, enabled: bool) -> Result<(), RestError> {
  let context = context.clone();

  let set_enabled_result = models::CurrencyModel::set_enabled(context, code, enabled).await;

  if set_enabled_result.is_err() {
    return Err(RestError::new("Error setting currency enabled".to_string()));
  }

  if set_enabled_result.unwrap() == 0 {
    return Err(RestError::new("Currency not found".to_string()));
  }

  Ok(())
}