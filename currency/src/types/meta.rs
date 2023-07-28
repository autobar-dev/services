use chrono::{DateTime, Utc};
use serde::Serialize;

#[derive(Debug, Clone)]
pub struct MetaFactors {
    pub hash: String,
    pub version: String,
    pub start_time: DateTime<Utc>,
}

#[derive(Debug, Serialize, Clone)]
pub struct Meta {
    pub hash: String,
    pub version: String,
    pub uptime: i64,
}
