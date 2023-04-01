use crate::models;
use crate::app_context;
use crate::types::RestError;

pub async fn new_currency_controller(context: app_context::Context, code: String, name: String, enabled: bool) -> Result<(), RestError> {
  let context = context.clone();
  let currency = models::CurrencyModel::new(context, code, name, enabled).await;

  if currency.is_err() {
    return Err(RestError::new("Error creating new currency".to_string()));
  }

  if currency.unwrap() == 0 {
    return Err(RestError::new("Error creating new currency".to_string()));
  }

  Ok(())
}