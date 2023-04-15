use chrono::{
    Utc,
    DateTime,
    serde::ts_seconds,
};
use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct Session {
    pub id: uuid::Uuid,
    pub user_id: i32,
    pub user_email: String,

    pub user_agent: Option<String>,

    #[serde(with = "ts_seconds")]
    pub valid_until: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub last_used: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub created_at: DateTime<Utc>,
}
