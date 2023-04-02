use crate::models;
use crate::types;
use crate::app_context;
use crate::types::RestError;

pub async fn get_remote_rate_controller(context: app_context::Context, from: String, to: String) -> Result<types::RemoteRate, RestError> {
  let context = context.clone();
  let rate = models::RemoteRateModel::get(context, from, to).await;

  if rate.is_err() {
    return Err(RestError::new("Error fetching rate".to_string()));
  }

  let rate = types::RemoteRate::from(rate.unwrap());

  Ok(rate)
}