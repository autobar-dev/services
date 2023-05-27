use chrono::{serde::ts_seconds, DateTime, Utc};
use serde::Serialize;

use crate::types;

#[derive(Clone, Serialize, Debug, sqlx::FromRow)]
pub struct ModuleModel {
    pub id: i32,
    pub serial_number: String,
    pub private_key: String,

    #[serde(with = "ts_seconds")]
    pub created_at: DateTime<Utc>,
}

impl ModuleModel {
    pub async fn get_by_serial_number(
        context: types::AppContext,
        serial_number: String,
    ) -> Result<ModuleModel, sqlx::Error> {
        let mut conn = context.database_pool.acquire().await?;

        sqlx::query_as!(
            ModuleModel,
            "SELECT *
            FROM modules
            WHERE serial_number = $1;",
            serial_number
        )
        .fetch_one(&mut conn)
        .await
    }

    pub async fn create(
        context: types::AppContext,
        serial_number: String,
        private_key_hash: String,
    ) -> Result<(), sqlx::Error> {
        let mut conn = context.database_pool.acquire().await?;

        let result = sqlx::query!(
            "INSERT INTO modules
            (serial_number, private_key)
            VALUES ($1, $2);",
            serial_number,
            private_key_hash
        )
        .execute(&mut conn)
        .await?;

        Ok(())
    }
}
