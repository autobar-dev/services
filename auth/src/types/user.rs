use crate::models;

use chrono::{DateTime, Utc};
use serde::Serialize;

#[derive(Clone, Serialize, Debug)]
pub struct User {
    pub id: i32,
    pub email: String,

    pub created_at: DateTime<Utc>,
}

impl User {
    pub fn from(user: models::UserModel) -> User {
        User {
            id: user.id,
            email: user.email,

            created_at: user.created_at,
        }
    }
}
