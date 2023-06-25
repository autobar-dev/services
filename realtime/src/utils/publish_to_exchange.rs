use lapin::{options::BasicPublishOptions, BasicProperties};
use serde_json::json;

use crate::{
    types::{AppContext, ClientType, Message},
    utils::client_identifier_to_exchange_name,
};

pub async fn publish_to_exchange(
    context: AppContext,
    client_type: ClientType,
    client_identifier: String,
    message: Message,
) -> anyhow::Result<()> {
    let message_string = json!(message).to_string();
    let bytes_to_publish = message_string.as_bytes();

    log::debug!("string message payload to publish: {}", message_string);

    context
        .amqp_channel
        .basic_publish(
            &client_identifier_to_exchange_name(client_type, client_identifier),
            "",
            BasicPublishOptions::default(),
            bytes_to_publish,
            BasicProperties::default(),
        )
        .await?;

    Ok(())
}
