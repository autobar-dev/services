use crate::models;

use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct RemoteRate {
    pub from: String,
    pub to: String,
    pub rate: f64,
}

impl RemoteRate {
    pub fn from(remote_rate: models::RemoteRateModel) -> RemoteRate {
        let from = remote_rate.base_code.to_uppercase();
        let to = remote_rate.target_code.to_uppercase();

        RemoteRate {
            from,
            to,
            rate: remote_rate.conversion_rate,
        }
    }
}

