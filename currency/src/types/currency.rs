use crate::models;

use chrono::{DateTime, Utc};
use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct Currency {
    pub id: i32,
    pub code: String,
    pub name: String,
    pub enabled: bool,
    pub updated_at: DateTime<Utc>,
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

