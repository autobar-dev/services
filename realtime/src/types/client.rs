use actix_web_lab::sse::{self, SendError};
use chrono::{DateTime, Utc};
use futures_lite::stream::StreamExt;
use lapin::options::{
    BasicAckOptions, BasicConsumeOptions, ExchangeDeclareOptions, QueueBindOptions,
    QueueDeclareOptions, QueueDeleteOptions,
};
use lapin::types::FieldTable;
use lapin::ExchangeKind;
use std::str;

use std::time::Duration;

use crate::types;
use crate::utils::client_identifier_to_exchange_name;

use super::Message;

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

    pub queue_name: Option<String>,
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
            queue_name: None,
        }
    }

    pub async fn send(self, message: String) -> Result<(), SendError> {
        log::debug!("raw message to send: {:?}", message);

        let body: String;
        let parsed_message: Message = serde_json::from_str(message.as_str()).unwrap();

        let event_name = match parsed_message {
            Message::Command(_) => {
                body = message.clone();
                "command"
            }
            Message::Simple(simple_message) => {
                body = simple_message.body;
                "simple"
            }
        };

        let send_result = self
            .sse_sender
            .send(sse::Data::new(body).event(event_name))
            .await;

        if send_result.is_err() {
            return Err(send_result.unwrap_err());
        }

        Ok(())
    }

    pub async fn listen(&mut self, context: types::AppContext) -> anyhow::Result<()> {
        let client_type = self.client_type;
        let identifier = self.clone().identifier.clone();

        let _ = context
            .amqp_channel
            .exchange_declare(
                &client_identifier_to_exchange_name(client_type, identifier.clone()),
                ExchangeKind::Fanout,
                ExchangeDeclareOptions::default(),
                FieldTable::default(),
            )
            .await?;

        let mut queue_declare_options = QueueDeclareOptions::default();
        queue_declare_options.exclusive = true;
        let queue = context
            .amqp_channel
            .queue_declare("", queue_declare_options, FieldTable::default())
            .await?;

        self.queue_name = Some(queue.name().to_string());

        let _ = context
            .amqp_channel
            .queue_bind(
                self.queue_name.clone().unwrap().as_str(),
                &client_identifier_to_exchange_name(client_type, identifier.clone()),
                "",
                QueueBindOptions::default(),
                FieldTable::default(),
            )
            .await?;

        let consumer_tag = uuid::Uuid::new_v4().simple().to_string();
        let mut consumer = context
            .amqp_channel
            .basic_consume(
                self.queue_name.clone().unwrap().as_str(),
                consumer_tag.as_str(),
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await?;

        log::info!(
            "consuming queue {} as consumer_tag {}",
            self.queue_name.clone().unwrap(),
            consumer_tag
        );

        self.state = ClientState::Listening;

        let mut heartbeat_interval = tokio::time::interval(Duration::from_secs_f32(
            context.config.sse_heartbeat_interval,
        ));

        heartbeat_interval.tick().await; // ticks immediately

        let heartbeat_context_clone = context.clone();
        let heartbeat_self_clone = self.clone();

        log::debug!(
            "listening on connection {} ({})",
            self.identifier,
            self.client_type
        );

        // heartbeat thread
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

        // queue consumer and client sender thread
        tokio::spawn(async move {
            while let Some(delivery) = consumer.next().await {
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
        });

        Ok(())
    }

    // pub async fn recovery(mut self, context: types::AppContext) -> anyhow::Result<()> {
    //     self.state = ClientState::Recovery;
    // }

    pub async fn cancel(mut self, context: types::AppContext) -> anyhow::Result<()> {
        self.state = ClientState::Cancelling;
        let queue_name = self.queue_name.clone().unwrap();

        log::debug!(
            "cancelling {} ({}) connection",
            self.identifier,
            self.client_type
        );

        context
            .amqp_channel
            .queue_delete(queue_name.as_str(), QueueDeleteOptions::default())
            .await?;

        log::debug!(
            "cancelled {} ({}) connection",
            self.identifier,
            self.client_type
        );

        Ok(())
    }
}
