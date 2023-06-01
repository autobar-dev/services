use actix_web::{post, web, HttpResponse, Responder};
use deadpool_redis::redis;

use serde::Deserialize;

use crate::{
    types::{AppContext, ClientType, CommandMessage, Message},
    utils::{client_identifier_to_redis_key, publish_to_queue},
};

#[derive(Clone, Debug, Deserialize)]
pub struct SendCommandBody {
    client_type: String,
    identifier: String,
    command: String,
    args: String,
}

#[post("/send-command")]
pub async fn send_command_route(
    data: web::Data<AppContext>,
    body: web::Json<SendCommandBody>,
) -> impl Responder {
    let context = data.as_ref().to_owned();

    if serde_json::from_str::<serde_json::Value>(body.args.as_str()).is_err() {
        return HttpResponse::BadRequest().body("command body cannot be parsed");
    }

    let client_type: Option<ClientType> = match body.client_type.to_lowercase().as_str() {
        "module" => Some(ClientType::Module),
        "user" => Some(ClientType::User),
        _ => None,
    };

    match client_type {
        None => {
            return HttpResponse::BadRequest().body("invalid client type ");
        }
        Some(ClientType::User) => {
            return HttpResponse::BadRequest()
                .body("invalid client type. only modules can consume commands");
        }
        Some(ClientType::Module) => {}
    };

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

    if !client_connected {
        return HttpResponse::NotFound().body("client not found");
    }

    let message = CommandMessage {
        command: body.command.clone(),
        args: body.args.clone(),
    };

    let publish_result = publish_to_queue(
        context,
        client_type,
        body.identifier.clone(),
        Message::Command(message),
    )
    .await;

    if publish_result.is_err() {
        return HttpResponse::InternalServerError().body("could not deliver message");
    }

    HttpResponse::Ok().body("successfully sent")
}
