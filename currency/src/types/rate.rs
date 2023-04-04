use crate::models;

use serde::Serialize;
use chrono::{
  DateTime,
  Utc,
  serde::ts_seconds,
};

#[derive(Debug, Serialize)]
pub struct Rate {
  pub from: String,
  pub to: String,
  pub rate: f64,
  
  #[serde(with = "ts_seconds")]
  pub updated_at: DateTime<Utc>,
}

impl Rate {
  pub fn from(rate: models::RateModel) -> Rate {
    Rate {
      from: rate.from,
      to: rate.to,
      rate: rate.rate,
      updated_at: rate.updated_at,
    }
  }
}