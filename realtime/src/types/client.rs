use actix_web_lab::sse::{self, SendError};
use chrono::{DateTime, Utc};
use deadpool_redis::redis;
use futures_lite::stream::StreamExt;
use lapin::options::{
    BasicAckOptions, BasicConsumeOptions, QueueDeclareOptions, QueueDeleteOptions,
};
use lapin::types::FieldTable;
use std::str;

use std::time::Duration;

use crate::types;
use crate::utils::{client_identifier_to_queue_name, client_identifier_to_redis_key};

#[derive(Debug, Clone, PartialEq)]
pub enum ClientState {
    Initializing,
    Listening,
    Cancelling,
}

#[derive(Clone, Debug)]
pub struct Client {
    pub state: ClientState,

    pub client_type: types::ClientType,
    pub identifier: String,
    pub sse_sender: sse::Sender, // shouldnt be pub?
}

impl Client {
    pub fn new(
        client_type: types::ClientType,
        identifier: String,
        sse_sender: sse::Sender,
    ) -> Self {
        Self {
            state: ClientState::Initializing,
            client_type,
            identifier,
            sse_sender,
        }
    }

    pub async fn send(self, message: String) -> Result<(), SendError> {
        let send_result = self
            .sse_sender
            .send(sse::Data::new(message).event("message"))
            .await;

        if send_result.is_err() {
            return Err(send_result.unwrap_err());
        }

        Ok(())
    }

    pub async fn listen(mut self, context: types::AppContext) -> anyhow::Result<()> {
        self.state = ClientState::Listening;

        let now: DateTime<Utc> = Utc::now();
        let now_string = now.to_rfc3339();

        let mut conn = context.clone().redis_pool.get().await?;

        let client_type = self.client_type;
        let identifier = self.clone().identifier;

        redis::cmd("SET")
            .arg(&[
                client_identifier_to_redis_key(client_type, identifier.clone()),
                now_string,
            ])
            .query_async(&mut conn)
            .await?;

        let mut heartbeat_interval = tokio::time::interval(Duration::from_secs_f32(
            context.config.sse_heartbeat_interval,
        ));

        heartbeat_interval.tick().await; // ticks immediately

        let heartbeat_context_clone = context.clone();

        let _ = context
            .amqp_channel
            .queue_declare(
                client_identifier_to_queue_name(client_type, identifier.clone()).as_str(),
                QueueDeclareOptions::default(),
                FieldTable::default(),
            )
            .await?;

        let consumer_tag = uuid::Uuid::new_v4().simple().to_string();
        let mut consumer = context
            .amqp_channel
            .basic_consume(
                client_identifier_to_queue_name(client_type, identifier.clone()).as_str(),
                consumer_tag.as_str(),
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await?;

        let heartbeat_self_clone = self.clone();

        tokio::spawn(async move {
            while heartbeat_self_clone.state == ClientState::Listening {
                let now: DateTime<Utc> = Utc::now();
                let now_string = now.to_rfc3339();

                let send_result = heartbeat_self_clone
                    .clone()
                    .sse_sender
                    .send(sse::Data::new(now_string).event("heartbeat"))
                    .await;

                if send_result.is_err() {
                    // go into recovery
                    heartbeat_self_clone
                        .cancel(heartbeat_context_clone)
                        .await
                        .unwrap_or_else(|err| log::error!("failed to cancel: {:?}", err));
                    break;
                }

                heartbeat_interval.tick().await;
            }
        });

        let consumer_self_clone = self.clone();

        tokio::spawn(async move {
            log::debug!("starting to listen on consumer");

            while let Some(delivery) = consumer.next().await {
                log::debug!("delivery consumed: {:?}", delivery);

                let loop_self_clone = consumer_self_clone.clone();

                if let Ok(delivery) = delivery {
                    let data = str::from_utf8(&delivery.data);

                    if data.is_err() {
                        log::error!(
                            "value from delivery is not valid UTF-8: {:?}",
                            data.unwrap_err()
                        );
                    } else {
                        let data = data.unwrap().to_string();

                        let send_result = loop_self_clone.clone().send(data).await;

                        if send_result.is_err() {
                            // go into recovery
                            loop_self_clone
                                .cancel(context.clone())
                                .await
                                .unwrap_or_else(|err| panic!("failed to cancel client: {:?}", err));
                            break;
                        }

                        delivery
                            .ack(BasicAckOptions::default())
                            .await
                            .expect("basick_ack failed");
                    }
                }
            }

            log::debug!("stopped listening on consumer");
        });

        Ok(())
    }

    // pub async fn recovery(mut self, context: types::AppContext) -> anyhow::Result<()> {
    //     self.state = ClientState::Recovery;
    // }

    pub async fn cancel(mut self, context: types::AppContext) -> anyhow::Result<()> {
        self.state = ClientState::Cancelling;

        let mut conn = context.redis_pool.get().await?;

        redis::cmd("DEL")
            .arg(client_identifier_to_redis_key(
                self.client_type,
                self.clone().identifier,
            ))
            .query_async(&mut conn)
            .await?;

        context
            .amqp_channel
            .queue_delete(
                &client_identifier_to_queue_name(self.client_type, self.clone().identifier),
                QueueDeleteOptions::default(),
            )
            .await?;

        Ok(())
    }
}
