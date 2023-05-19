use crate::types;

use actix_web_lab::sse;
use anyhow;
use chrono::{DateTime, Utc};
use deadpool_redis::redis;

pub struct ModuleConnection {
    pub serial_number: String,
    pub sse_sender: sse::Sender, // shouldnt be pub?
}

impl ModuleConnection {
    pub async fn new(serial_number: String, sse_sender: sse::Sender) -> Self {
        Self {
            serial_number,
            sse_sender,
        }
    }

    pub async fn init(self, context: types::AppContext) -> anyhow::Result<()> {
        let mut conn = context.redis_pool.get().await?;

        let now: DateTime<Utc> = Utc::now();
        let now_string = now.to_rfc3339();

        // TODO: Initialize pub/sub from a message queue on topic SERIAL_NUMBER
        let mut amqp_channel = context.amqp_channel.queue_declare(, options, arguments)

        redis::cmd("SET")
            .arg(&[format!("connected:{}", self.serial_number), now_string])
            .query_async(&mut conn)
            .await?;

        Ok(())
    }

    pub async fn send(self, )

    pub async fn on_disconnect(self, context: types::AppContext) -> anyhow::Result<()> {
        // shouldnt
        // be pub ?
        let mut conn = context.redis_pool.get().await?;

        redis::cmd("DEL")
            .arg(format!("connected:{}", self.serial_number))
            .query_async(&mut conn)
            .await?;

        Ok(())
    }
}
