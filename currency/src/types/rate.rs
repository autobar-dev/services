use crate::models;

use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;

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
            from: rate.from_currency,
            to: rate.to_currency,
            rate: rate.rate,
            updated_at: rate.updated_at,
        }
    }
}

