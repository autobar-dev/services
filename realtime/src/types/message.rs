use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct SimpleMessage {
    pub body: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct CommandMessage {
    pub command: String,
    pub body: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
#[serde(untagged)]
pub enum Message {
    Command(CommandMessage),
    Simple(SimpleMessage),
}
