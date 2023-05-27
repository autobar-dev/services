use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;
use uuid::{self, Uuid};

use crate::types::{self, ClientType};

#[derive(Clone, Serialize, Debug, sqlx::FromRow)]
pub struct SessionModel {
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

#[derive(Debug, Clone, sqlx::Type)]
struct ClientTypeRecord {
    client_type: ClientType,
}

impl SessionModel {
    pub async fn get_by_session_id(
        context: types::AppContext,
        session_uuid: Uuid,
    ) -> Result<SessionModel, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let client_type: ClientTypeRecord = sqlx::query_as!(
            ClientTypeRecord,
            r#"SELECT client_type AS "client_type: ClientType"
            FROM sessions s
            WHERE s.id = $1;"#,
            session_uuid
        )
        .fetch_one(&mut conn)
        .await?;

        match client_type.client_type {
            ClientType::User => {
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT
                    s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type as "client_type: ClientType", u.email AS client_identifier
                    FROM sessions s
                    INNER JOIN users u
                    ON s.client_identifier = u.id
                    WHERE s.id = $1;"#,
                    session_uuid
                )
                .fetch_one(&mut conn)
                .await
            },
            ClientType::Module => {
                
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT
                    s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type as "client_type: ClientType", m.serial_number AS client_identifier
                    FROM sessions s
                    INNER JOIN modules m
                    ON s.client_identifier = m.id
                    WHERE s.id = $1;"#,
                    session_uuid
                )
                .fetch_one(&mut conn)
                .await
            },
        }
    }

    pub async fn get_by_internal_id(
        context: types::AppContext,
        internal_id: i32,
    ) -> Result<SessionModel, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let client_type = sqlx::query_as!(
            ClientTypeRecord,
            r#"SELECT client_type AS "client_type: ClientType"
            FROM sessions s
            WHERE s.internal_id = $1;"#,
            internal_id
        )
        .fetch_one(&mut conn)
        .await?;

        let session_data: Result<SessionModel, sqlx::Error> = match client_type.client_type {
            ClientType::User => {
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type as "client_type: ClientType", u.email AS client_identifier
                    FROM sessions s
                    INNER JOIN users u
                    ON s.client_identifier = u.id
                    WHERE s.internal_id = $1;"#,
                    internal_id
                )
                .fetch_one(&mut conn)
                .await
            }
            ClientType::Module => {
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type AS "client_type: ClientType", m.serial_number AS client_identifier
                    FROM sessions s
                    INNER JOIN modules m
                    ON s.client_identifier = m.id
                    WHERE s.internal_id = $1;"#,
                    internal_id
                )
                .fetch_one(&mut conn)
                .await
            }
        };

        session_data
    }

    pub async fn all_for_client(
        context: types::AppContext,
        client_type: ClientType,
        client_identifier: String,
    ) -> Result<Vec<SessionModel>, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        match client_type {
            ClientType::User => {
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type AS "client_type: ClientType", u.email AS client_identifier
                    FROM sessions s
                    INNER JOIN users u
                    ON s.client_identifier = u.id
                    WHERE u.email = $1;"#,
                    client_identifier
                )
                .fetch_all(&mut conn)
                .await
            },
            ClientType::Module => {
                sqlx::query_as!(
                    SessionModel,
                    r#"SELECT s.id, s.user_agent, s.valid_until, s.last_used, s.created_at, s.internal_id, s.client_type AS "client_type: ClientType", m.serial_number AS client_identifier
                    FROM sessions s
                    INNER JOIN modules m
                    ON s.client_identifier = m.id
                    WHERE m.serial_number = $1;"#,
                    client_identifier
                )
                .fetch_all(&mut conn)
                .await
            },
        }
    }

    pub async fn create(
        context: types::AppContext,
        client_type: ClientType,
        client_identifier: String,
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

        match client_type {
            ClientType::User => {
                let result = sqlx::query!(
                    r#"INSERT INTO sessions
                    (client_identifier, user_agent, valid_until, client_type)
                    SELECT u.id, $1, $2, 'user'
                    FROM users u
                    WHERE u.email = $3
                    RETURNING id;"#,
                    user_agent,
                    valid_until,
                    client_identifier
                )
                .fetch_one(&mut conn)
                .await?;

                Ok(result.id)
            },
            ClientType::Module => {
                let result = sqlx::query!(
                    r#"INSERT INTO sessions
                    (client_identifier, user_agent, valid_until, client_type)
                    SELECT m.id, $1, $2, 'module'
                    FROM modules m
                    WHERE m.serial_number = $3
                    RETURNING id;"#,
                    user_agent,
                    valid_until,
                    client_identifier
                )
                .fetch_one(&mut conn)
                .await?;

                Ok(result.id)
            }
        }
    }

    pub async fn update_last_used(
        context: types::AppContext,
        session_uuid: Uuid,
    ) -> Result<u64, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query!(
            "UPDATE sessions
            SET last_used = CURRENT_TIMESTAMP
            WHERE id = $1;",
            session_uuid
        )
        .execute(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error updating last_used: {:?}", result_err);
            return Err(result_err);
        }

        let result = result.unwrap();

        Ok(result.rows_affected())
    }

    pub async fn delete_all_expired_for_client(
        context: types::AppContext,
        client_type: ClientType,
        client_identifier: String,
    ) -> Result<u64, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = match client_type {
            ClientType::User => {
                sqlx::query!(
                    "DELETE FROM sessions s
                    USING users u
                    WHERE s.client_identifier = u.id
                    AND u.email = $1
                    AND s.valid_until < CURRENT_TIMESTAMP;",
                    client_identifier
                )
                .execute(&mut conn)
                .await?
            },
            ClientType::Module => {
                sqlx::query!(
                    "DELETE FROM sessions s
                    USING modules m
                    WHERE s.client_identifier = m.id
                    AND m.serial_number = $1
                    AND s.valid_until < CURRENT_TIMESTAMP;",
                    client_identifier
                )
                .execute(&mut conn)
                .await?
            },
        };

        Ok(result.rows_affected())
    }

    pub async fn delete(
        context: types::AppContext,
        session_uuid: Uuid,
    ) -> Result<u64, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query!(
            "DELETE FROM sessions
            WHERE id = $1;",
            session_uuid
        )
        .execute(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error deleting session: {:?}", result_err);
            return Err(result_err);
        }

        let result = result.unwrap();

        Ok(result.rows_affected())
    }
}
