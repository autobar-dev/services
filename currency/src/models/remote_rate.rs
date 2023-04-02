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
pub struct RemoteRateErrorModel {
  pub result: String,
  pub documentation: String,

  #[serde(rename = "terms-of-use")] // yes, they use kebab case only in the error response
  pub terms_of_use: String,

  #[serde(rename = "error-type")]
  pub error_type: String,
}

impl RemoteRateModel {
  pub async fn get(context: Context, from: String, to: String) -> Result<RemoteRateModel, RemoteRateErrorModel> {
    let url = format!("https://v6.exchangerate-api.com/v6/{}/pair/{}/{}", context.config.exchange_rate_api_key, from, to);
    let response = reqwest::get(&url).await;

    if response.is_err() {
      log::error!("Error while fetching remote rate: {}", response.unwrap_err());

      return Err(RemoteRateErrorModel {
        result: "error".to_string(),
        documentation: "https://www.exchangerate-api.com/docs".to_string(),
        terms_of_use: "https://www.exchangerate-api.com/terms".to_string(),
        error_type: "unknown".to_string(),
      });
    }

    let response = response.unwrap();
    let body = response.text().await;

    if body.is_err() {
      log::error!("Error while fetching remote rate: {}", body.unwrap_err());

      return Err(RemoteRateErrorModel {
        result: "error".to_string(),
        documentation: "https://www.exchangerate-api.com/docs".to_string(),
        terms_of_use: "https://www.exchangerate-api.com/terms".to_string(),
        error_type: "unknown".to_string(),
      });
    }

    let body = body.unwrap();

    let rate: Result<RemoteRateModel, _> = serde_json::from_str(&body);

    if rate.is_err() {
      log::error!("Error while fetching remote rate: {}", rate.unwrap_err());

      let error: Result<RemoteRateErrorModel, _> = serde_json::from_str(&body);

      if error.is_err() {
        log::error!("Error while fetching remote rate: {}", error.unwrap_err());

        return Err(RemoteRateErrorModel {
          result: "error".to_string(),
          documentation: "https://www.exchangerate-api.com/docs".to_string(),
          terms_of_use: "https://www.exchangerate-api.com/terms".to_string(),
          error_type: "unknown".to_string(),
        });
      }

      return Err(error.unwrap());
    }

    let rate = rate.unwrap();

    Ok(rate)
  }
}

// {
//     "result": "error",
//     "documentation": "https://www.exchangerate-api.com/docs",
//     "terms-of-use": "https://www.exchangerate-api.com/terms",
//     "error-type": "unsupported-code"
// }

// {
//     "result": "error",
//     "documentation": "https://www.exchangerate-api.com/docs",
//     "terms-of-use": "https://www.exchangerate-api.com/terms",
//     "error-type": "invalid-key"
// }