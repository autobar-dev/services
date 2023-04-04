use crate::app_context::Context;

use serde::{
  Serialize,
  Deserialize,
};

#[derive(Debug, Serialize, Deserialize)]
pub struct RemoteRateModel {
  pub result: String,
  pub documentation: String,
  pub terms_of_use: String,
  pub time_last_update_unix: u32,
  pub time_last_update_utc: String,
  pub time_next_update_unix: u32,
  pub time_next_update_utc: String,
  pub base_code: String,
  pub target_code: String,
  pub conversion_rate: f64,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct RemoteRateApiError {
  pub result: String,
  pub documentation: String,

  #[serde(rename = "terms-of-use")] // yes, they use kebab case only in the error response
  pub terms_of_use: String,

  #[serde(rename = "error-type")]
  pub error_type: String,
}

#[derive(Debug, Serialize)]
pub struct RemoteRateError {
  pub message: String,
}

impl RemoteRateModel {
  pub async fn get(context: Context, from: String, to: String) -> Result<RemoteRateModel, RemoteRateError> {
    let url = format!(
      "https://v6.exchangerate-api.com/v6/{}/pair/{}/{}",
      context.config.exchange_rate_api_key,
      from,
      to
    );
    let response = reqwest::get(&url).await;

    if response.is_err() {
      log::error!("Error making request to remote exchange API: {}", response.unwrap_err());

      return Err(RemoteRateError {
        message: "communication with remote exchange API failed".to_string(),
      });
    }

    let response = response.unwrap();
    let body = response.text().await;

    if body.is_err() {
      return Err(RemoteRateError {
        message: "communication with remote exchange API failed".to_string(),
      });
    }

    let body = body.unwrap();

    let rate_success: Result<RemoteRateModel, _> = serde_json::from_str(&body);
    let rate_error: Result<RemoteRateApiError, _> = serde_json::from_str(&body);

    if rate_success.is_ok() {
      let rate = rate_success.unwrap();
      return Ok(rate);
    }

    if rate_error.is_ok() {
      let rate_error = rate_error.unwrap();
      let rate_error_type = rate_error.error_type.as_str();

      match rate_error_type {
        "unsupported-code" => {
          return Err(RemoteRateError {
            message: "unsupported currency code".to_string(),
          });
        },
        "invalid-key" => {
          log::error!("Invalid API key for remote exchange API");

          return Err(RemoteRateError {
            message: "unable to fetch from remote API".to_string(),
          });
        },
        _ => {
          return Err(RemoteRateError {
            message: "unknown error".to_string(),
          });
        }
      }
    }    

    Err(RemoteRateError {
      message: "unknown error".to_string(),
    })
  }
}
