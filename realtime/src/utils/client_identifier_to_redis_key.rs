use crate::types::ClientType;

pub fn client_identifier_to_redis_key(client_type: ClientType, identifier: String) -> String {
    match client_type {
        ClientType::Module => format!("connected:module:{}", identifier),
        ClientType::User => format!("connected:user:{}", identifier),
    }
}
