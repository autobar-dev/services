use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;

use crate::types;

#[derive(Clone, Serialize, Debug, sqlx::FromRow)]
pub struct UserModel {
    pub id: i32,
    pub email: String,
    pub password: String,

    #[serde(with = "ts_seconds")]
    pub created_at: DateTime<Utc>,
}

impl UserModel {
    pub async fn _get_all(context: types::AppContext) -> Result<Vec<UserModel>, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query_as!(
            UserModel,
            "SELECT *
            FROM users;"
        )
        .fetch_all(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error getting all users: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap())
    }

    pub async fn get_by_email(context: types::AppContext, email: String) -> Result<UserModel, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query_as!(
            UserModel,
            "SELECT *
            FROM users
            WHERE email = $1;",
            email
        )
        .fetch_one(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error getting user by email: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap())
    }
}
