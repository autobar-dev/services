use crate::models;
use crate::types;
use crate::app_context;
use crate::types::RestError;

pub async fn get_currency_controller(context: app_context::Context, code: String) -> Result<types::Currency, RestError> {
  let context = context.clone();
  let currency = models::CurrencyModel::get(context, code).await;

  if currency.is_err() {
    return Err(RestError::new("Error fetching currencies".to_string()));
  }

  let currency = types::Currency::from(currency.unwrap());

  Ok(currency)
}