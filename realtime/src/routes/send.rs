use actix_web::{post, web, HttpResponse, Responder};

use serde::Deserialize;

use crate::{
    types::{AppContext, ClientType, Message, SimpleMessage},
    utils::publish_to_exchange,
};

#[derive(Clone, Debug, Deserialize)]
pub struct SendBody {
    client_type: String,
    identifier: String,
    body: String,
}

#[post("/send")]
pub async fn send_route(data: web::Data<AppContext>, body: web::Json<SendBody>) -> impl Responder {
    let context = data.as_ref().to_owned();

    let client_type: Option<ClientType> = match body.client_type.to_lowercase().as_str() {
        "module" => Some(ClientType::Module),
        "user" => Some(ClientType::User),
        _ => None,
    };

    if client_type.is_none() {
        return HttpResponse::BadRequest().body("incorrect client type");
    }

    let client_type = client_type.unwrap();

    let message = SimpleMessage {
        body: body.body.clone(),
    };

    let publish_result = publish_to_exchange(
        context,
        client_type,
        body.identifier.clone(),
        Message::Simple(message),
    )
    .await;

    if publish_result.is_err() {
        return HttpResponse::InternalServerError().body("could not deliver message");
    }

    HttpResponse::Ok().body("successfully submitted")
}
