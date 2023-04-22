use crate::app_context::Context;

use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, sqlx::FromRow)]
pub struct RateModel {
    pub id: i32,
    pub from_currency: String,
    pub to_currency: String,
    pub rate: f64,

    #[serde(with = "ts_seconds")]
    pub updated_at: DateTime<Utc>,
}

impl RateModel {
    pub async fn get(context: Context, from: String, to: String) -> Result<RateModel, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query_as!(
            RateModel,
            "SELECT * FROM rates WHERE from_currency = $1 AND to_currency = $2;",
            from,
            to
        )
        .fetch_one(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error fetching rate: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap())
    }

    pub async fn set(
        context: Context,
        from: String,
        to: String,
        rate: f64,
    ) -> Result<u64, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query!(
            "
            INSERT INTO rates
            (from_currency, to_currency, rate)
            VALUES ($1, $2, $3)
            ON CONFLICT
            ON CONSTRAINT rates_from_currency_to_currency_key
            DO UPDATE SET
              rate = $3,
              updated_at = CURRENT_TIMESTAMP;
            ",
            from,
            to,
            rate
        )
        .execute(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error setting rate: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap().rows_affected())
    }

    pub async fn delete(context: Context, from: String, to: String) -> Result<u64, sqlx::Error> {
        let conn = context.database_pool.acquire().await;

        if conn.is_err() {
            let conn_err = conn.unwrap_err();

            log::error!("Error acquiring connection: {:?}", conn_err);
            return Err(conn_err);
        }

        let mut conn = conn.unwrap();

        let result = sqlx::query!(
            "DELETE FROM rates WHERE from_currency = $1 AND to_currency = $2;",
            from,
            to
        )
        .execute(&mut conn)
        .await;

        if result.is_err() {
            let result_err = result.unwrap_err();

            log::error!("Error deleting currency: {:?}", result_err);
            return Err(result_err);
        }

        Ok(result.unwrap().rows_affected())
    }
}
