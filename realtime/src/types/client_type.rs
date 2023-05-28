use std::fmt;

use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, Serialize, Deserialize)]
#[serde(rename_all = "lowercase")]
pub enum ClientType {
    Module,
    User,
}

impl fmt::Display for ClientType {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Self::Module => write!(f, "module"),
            Self::User => write!(f, "user"),
        }
    }
}
