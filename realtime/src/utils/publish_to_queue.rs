use lapin::{options::BasicPublishOptions, BasicProperties};
use serde_json::json;

use crate::types::{AppContext, ClientType, Message};

use super::client_identifier_to_queue_name;

pub async fn publish_to_queue(
    context: AppContext,
    client_type: ClientType,
    client_identifier: String,
    message: Message,
) -> anyhow::Result<()> {
    context
        .amqp_channel
        .basic_publish(
            "",
            client_identifier_to_queue_name(client_type, client_identifier).as_str(),
            BasicPublishOptions::default(),
            json!(message).to_string().as_bytes(),
            BasicProperties::default(),
        )
        .await?;

    Ok(())
}
