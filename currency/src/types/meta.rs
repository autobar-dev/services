use serde::Serialize;

#[derive(Debug, Serialize, Clone)]
pub struct Meta {
  pub hash: String,
  pub version: String,
}