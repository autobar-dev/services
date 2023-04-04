use crate::models;
use crate::app_context;
use crate::types::RestError;

pub async fn set_rate_controller(context: app_context::Context, from: String, to: String, rate: f64) -> Result<(), RestError> {
  let context = context.clone();

  let set_rate_result = models::RateModel::set(context, from, to, rate).await;

  if set_rate_result.is_err() {
    return Err(RestError::new("Error setting rate".to_string()));
  }

  if set_rate_result.unwrap() == 0 {
    return Err(RestError::new("Error setting rate".to_string()));
  }

  Ok(())
}