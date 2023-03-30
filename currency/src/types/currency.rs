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
  pub fn from(enabled_currency: models::CurrencyModel) -> Currency {
    Currency {
      id: enabled_currency.id,
      code: enabled_currency.code,
      name: enabled_currency.name,
      enabled: enabled_currency.enabled,
      updated_at: enabled_currency.updated_at,
      created_at: enabled_currency.created_at,
    }
  }
}