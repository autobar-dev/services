use std::fmt;

use serde::Serialize;

#[derive(Debug, Clone, Copy, Serialize)]
pub enum ClientType {
    Module,
    User,
}

impl fmt::Display for ClientType {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Self::Module => write!(f, "MODULE"),
            Self::User => write!(f, "USER"),
        }
    }
}
