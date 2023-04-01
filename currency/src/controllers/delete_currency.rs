use crate::models;
use crate::app_context;
use crate::types::RestError;

pub async fn delete_currency_controller(context: app_context::Context, code: String) -> Result<(), RestError> {
  let context = context.clone();

  let delete_result = models::CurrencyModel::delete(context, code).await;

  if delete_result.is_err() {
    return Err(RestError::new("Error deleting currency".to_string()));
  }

  if delete_result.unwrap() == 0 {
    return Err(RestError::new("Currency not found".to_string()));
  }

  Ok(())
}