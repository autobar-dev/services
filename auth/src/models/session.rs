use serde::Serialize;
use chrono::{
    serde::ts_seconds,
    DateTime,
    Utc,
};
use uuid::{self, Uuid};

use crate::types;

#[derive(Clone, Serialize, Debug, sqlx::FromRow)]
pub struct SessionModel {
    pub id: uuid::Uuid,
    pub user_id: i32,
    pub user_agent: Option<String>,

    #[serde(with = "ts_seconds")]
    pub valid_until: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub last_used: DateTime<Utc>,

    #[serde(with = "ts_seconds")]
    pub created_at: DateTime<Utc>,
}

impl SessionModel {
    pub async fn get_by_session_id(context: types::AppContext, session_uuid: Uuid) -> Result<SessionModel, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query_as!(
            SessionModel,
            "SELECT *
            FROM sessions
            WHERE id = $1;",
            session_uuid 
        )
            .fetch_one(&mut conn)
            .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error fetching session: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap())
    }

    pub async fn new(
        context: types::AppContext,
        user_id: i32,
        user_agent: Option<String>,
        valid_until: DateTime<Utc>,
    ) -> Result<Uuid, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query!(
            "INSERT INTO sessions
            (user_id, user_agent, valid_until)
            VALUES ($1, $2, $3)
            RETURNING id;",
            user_id,
            user_agent,
            valid_until
        )
            .fetch_one(&mut conn)
            .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error creating new session: {:?}", result_err);
            return Err(result_err);
        }

        let result = result.unwrap();

        Ok(result.id)
    }
}
