use crate::models;

use serde::Serialize;
use chrono::{
  DateTime,
  Utc,
};

#[derive(Debug, Serialize)]
pub struct Currency {
  pub id: i32,
  pub code: String,
  pub name: String,
  pub enabled: bool,

  #[serde(with = "chrono::serde::ts_seconds")]
  pub updated_at: DateTime<Utc>,

  #[serde(with = "chrono::serde::ts_seconds")]
  pub created_at: DateTime<Utc>,
}

impl Currency {
  pub fn from(currency: models::CurrencyModel) -> Currency {
    Currency {
      id: currency.id,
      code: currency.code,
      name: currency.name,
      enabled: currency.enabled,
      updated_at: currency.updated_at,
      created_at: currency.created_at,
    }
  }
}