use chrono::{
  DateTime,
  Utc,
  serde::ts_seconds,
};
use serde::{Serialize};

use crate::app_context::Context;

#[derive(sqlx::FromRow, Debug, Serialize)]
pub struct CurrencyModel {
  pub id: i32,
  pub code: String,
  pub name: String,
  pub enabled: bool,

  #[serde(with = "ts_seconds")]
  pub created_at: DateTime<Utc>,

  #[serde(with = "ts_seconds")]
  pub updated_at: DateTime<Utc>,
}

impl CurrencyModel {
  pub async fn get_all(context: Context) -> Result<Vec<CurrencyModel>, sqlx::Error> {
    let conn = context.database_pool.acquire().await;

    if conn.is_err() {
      let conn_err = conn.unwrap_err();

      log::error!("Error acquiring connection: {:?}", conn_err);
      return Err(conn_err);
    }

    let mut conn = conn.unwrap();

    let result = sqlx::query_as::<_, CurrencyModel>("
      SELECT id, code, name, enabled, created_at, updated_at
      FROM supported_currencies;
    ")
      .fetch_all(&mut conn)
      .await;
      
    if result.is_err() {
      let result_err = result.unwrap_err();

      log::error!("Error fetching currencies: {:?}", result_err);
      return Err(result_err);
    }

    let rows = result.unwrap();

    Ok(rows)
  }

  pub async fn get_all_enabled(context: Context) -> Result<Vec<CurrencyModel>, sqlx::Error> {
    let conn = context.database_pool.acquire().await;

    if conn.is_err() {
      let conn_err = conn.unwrap_err();

      log::error!("Error acquiring connection: {:?}", conn_err);
      return Err(conn_err);
    }

    let mut conn = conn.unwrap();

    let result = sqlx::query_as::<_, CurrencyModel>("
      SELECT id, code, name, enabled, created_at, updated_at
      FROM supported_currencies
      WHERE enabled = true;
    ")
      .fetch_all(&mut conn)
      .await;
      
    if result.is_err() {
      let result_err = result.unwrap_err();

      log::error!("Error fetching currencies: {:?}", result_err);
      return Err(result_err);
    }

    let rows = result.unwrap();

    Ok(rows)
  }
}