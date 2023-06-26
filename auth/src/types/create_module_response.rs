use crate::models;

use chrono::{DateTime, Utc};
use serde::Serialize;

#[derive(Clone, Serialize, Debug)]
pub struct CreateModuleResponse {
    pub id: i32,
    pub serial_number: String,
    pub private_key: String,
    pub created_at: DateTime<Utc>,
}

impl CreateModuleResponse {
    pub fn from(module: models::ModuleModel, private_key: String) -> CreateModuleResponse {
        CreateModuleResponse {
            id: module.id,
            serial_number: module.serial_number,
            created_at: module.created_at,
            private_key,
        }
    }
}
