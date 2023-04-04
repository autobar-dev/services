use crate::models;
use crate::types;
use crate::app_context;

pub async fn get_rate_controller(context: app_context::Context, from: String, to: String, no_fetch: bool) -> Result<types::Rate, types::RestError> {
  let context = context.to_owned();
  
  let context_rate_clone = context.clone();
  let from_rate_clone = from.clone();
  let to_rate_clone = to.clone();

  // Fetch rate from database
  let rate = models::RateModel::get(context_rate_clone, from_rate_clone, to_rate_clone).await;

  let from_rate_clone = from.clone();
  let to_rate_clone = to.clone();

  if rate.is_err() {
    let context = context.clone();
    let remote_rate = models::RemoteRateModel::get(context, from_rate_clone, to_rate_clone).await;

    if remote_rate.is_err() {
      return Err(types::RestError::new("Error fetching rate".to_string())); // NotFound
    }
  }

  let mut rate = types::Rate::from(rate.unwrap());

  if no_fetch {
    return Ok(rate);
  }

  // Get from and to currencies
  let context_from_clone = context.clone();
  let context_to_clone = context.clone();
  let from_currency_clone = from.clone();
  let to_currency_clone = to.clone();

  let from_currency = models::CurrencyModel::get(context_from_clone, from_currency_clone).await;
  let to_currency = models::CurrencyModel::get(context_to_clone, to_currency_clone).await;

  if from_currency.is_err() || to_currency.is_err() {
    return Err(types::RestError::new("Error fetching currency".to_string())); // NotFound
  }

  let from_currency = from_currency.unwrap();
  let to_currency = to_currency.unwrap();

  // Check if the rate is old enough to refetch
  let now = chrono::Utc::now();
  let rate_is_old = rate.updated_at < now - chrono::Duration::seconds(context.config.past_rate_retention.into());

  if rate_is_old {
    // Fetch rate from remote
    let context_remote_rate_clone = context.clone();
    let remote_rate = models::RemoteRateModel::get(context_remote_rate_clone, from_currency.code, to_currency.code).await;

    if remote_rate.is_err() {
      return Err(types::RestError::new("Error fetching remote rate".to_string())); // NotFound
    }

    let remote_rate = remote_rate.unwrap();

    // Update rate in database
    let context_update_rate_clone = context.clone();
    let from_update_rate_clone = from.clone();
    let to_update_rate_clone = to.clone();

    let update_rate_result = models::RateModel::set(context_update_rate_clone, from_update_rate_clone, to_update_rate_clone, remote_rate.conversion_rate).await;

    if update_rate_result.is_err() {
      return Err(types::RestError::new("Error updating rate".to_string())); // NotFound
    }

    let new_rate = models::RateModel::get(context, from, to).await;

    if new_rate.is_err() {
      return Err(types::RestError::new("Error fetching rate".to_string())); // NotFound
    }

    rate = types::Rate::from(new_rate.unwrap());
  }

  Ok(rate)
}