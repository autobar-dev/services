use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct RestError {
  pub message: String,
}

impl RestError {
  pub fn new(message: String) -> RestError {
    RestError { message }
  }
}