use actix_web::{post, web, HttpResponse, Responder};
use deadpool_redis::redis::{self, AsyncCommands};
use lapin::{options::BasicPublishOptions, BasicProperties};
use serde::Deserialize;
use std::sync::Arc;

use crate::{
    types::{AppContext, ClientType},
    utils::{client_identifier_to_queue_name, client_identifier_to_redis_key},
};

#[derive(Clone, Debug, Deserialize)]
pub struct SendBody {
    client_type: String,
    identifier: String,
    body: String,
}

#[derive(Clone, Debug)]
struct Message {
    client_type: ClientType,
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

    let conn = context.redis_pool.get().await;

    if conn.is_err() {
        return HttpResponse::InternalServerError().body("could not get redis connection");
    }

    let mut conn = conn.unwrap();

    let client_connected: Result<bool, redis::RedisError> = redis::cmd("EXISTS")
        .arg(client_identifier_to_redis_key(
            client_type,
            body.identifier.clone(),
        ))
        .query_async(&mut conn)
        .await;

    if client_connected.is_err() {
        log::error!(
            "failed to retrieve if client is connected: {:?}",
            client_connected.unwrap_err()
        );

        return HttpResponse::InternalServerError()
            .body("failed to retrieve if client is connected");
    }

    let client_connected = client_connected.unwrap();

    if client_connected == false {
        return HttpResponse::NotFound().body("client not found");
    }

    let message = Message {
        client_type,
        identifier: body.identifier.clone(),
        body: body.body.clone(),
    };

    // let clients = Arc::clone(&context.clients).lock().unwrap().clone();
    // let client = clients.get(&message.identifier);
    //
    // if client.is_none() {}
    //
    // let client = client.unwrap().clone();

    let publish_result = context
        .amqp_channel
        .basic_publish(
            "",
            client_identifier_to_queue_name(message.client_type, message.identifier).as_str(),
            BasicPublishOptions::default(),
            message.body.as_bytes(),
            BasicProperties::default(),
        )
        .await;

    if publish_result.is_err() {
        return HttpResponse::InternalServerError().body("could not deliver message");
    }

    HttpResponse::Ok().body("successfully sent")
}
