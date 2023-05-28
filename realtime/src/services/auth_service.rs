use reqwest::header;
use serde::Deserialize;

use crate::types::{
    consts::{INTERNAL_HEADER_NAME, SERVICE_NAME},
    ClientType,
};

#[derive(Clone, Debug)]
pub struct AuthService {
    http_client: reqwest::Client,
    base_url: String,
}

#[derive(Deserialize, Clone, Debug)]
pub struct AuthServiceVerifyResponseData {
    pub client_type: ClientType,
    pub client_identifier: String,
}

#[derive(Deserialize, Clone, Debug)]
pub struct AuthServiceVerifyResponse {
    pub status: String,
    pub error: Option<String>,
    pub data: Option<AuthServiceVerifyResponseData>,
}

impl AuthService {
    pub fn new(base_url: String) -> Self {
        let mut headers = header::HeaderMap::new();
        headers.insert(
            INTERNAL_HEADER_NAME,
            header::HeaderValue::from_static(SERVICE_NAME),
        );

        let client = reqwest::Client::builder()
            .default_headers(headers)
            .build()
            .unwrap();

        Self {
            base_url,
            http_client: client,
        }
    }

    // return either None if session is not valid or Some(identifier) if valid
    pub async fn verify_session(
        self,
        session: String,
    ) -> anyhow::Result<Option<AuthServiceVerifyResponseData>> {
        let endpoint_url = format!("{}/session/verify", self.base_url);
        let body = self
            .http_client
            .get(endpoint_url)
            .query(&[("session_id", session)])
            .send()
            .await?
            .json::<AuthServiceVerifyResponse>()
            .await?;

        log::debug!("auth service response: {:?}", body);

        Ok(match body.data.is_some() {
            true => Some(body.data.unwrap()),
            false => None,
        })
    }
}
