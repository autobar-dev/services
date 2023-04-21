use crate::models;

use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct SessionInfo {
    pub internal_id: i32,

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

impl SessionInfo {
    pub fn from(session: models::SessionModel) -> SessionInfo {
        SessionInfo {
            internal_id: session.internal_id,

            user_id: session.user_id,
            user_email: session.user_email,

            user_agent: session.user_agent,

            valid_until: session.valid_until,
            last_used: session.last_used,
            created_at: session.created_at,
        }
    }
}
