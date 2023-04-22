use crate::models;

use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct EnabledCurrency {
  pub code: String,
  pub name: String,
}

impl EnabledCurrency {
  pub fn from(enabled_currency: models::CurrencyModel) -> EnabledCurrency {
    EnabledCurrency {
      code: enabled_currency.code.to_string(),
      name: enabled_currency.name,
    }
  }
}