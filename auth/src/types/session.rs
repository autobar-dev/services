use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;

use crate::models;

use super::ClientType;

#[derive(Debug, Serialize)]
pub struct Session {
    pub id: uuid::Uuid,
    pub internal_id: i32,

    pub client_type: ClientType,
    pub client_identifier: String,

    pub user_agent: Option<String>,

    #[serde(with = "ts_seconds")]
    pub valid_until: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub last_used: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub created_at: DateTime<Utc>,
}

impl Session {
    pub fn from(session: models::SessionModel) -> Session {
        Session {
            id: session.id,
            internal_id: session.internal_id,

            client_type: session.client_type,
            client_identifier: session.client_identifier,

            user_agent: session.user_agent,

            valid_until: session.valid_until,
            last_used: session.last_used,
            created_at: session.created_at,
        }
    }
}
