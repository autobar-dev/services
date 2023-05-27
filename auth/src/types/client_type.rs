use serde::Serialize;

#[derive(Clone, Copy, Debug, Serialize, sqlx::Type)]
#[sqlx(type_name = "client_type")]
#[sqlx(rename_all = "lowercase")]
pub enum ClientType {
    #[serde(rename = "module")]
    Module,

    #[serde(rename = "user")]
    User,
}
