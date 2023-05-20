use crate::types::ClientType;

pub fn client_identifier_to_queue_name(client_type: ClientType, identifier: String) -> String {
    match client_type {
        ClientType::Module => format!("mod_{}", identifier),
        ClientType::User => format!("user_{}", identifier),
    }
}
