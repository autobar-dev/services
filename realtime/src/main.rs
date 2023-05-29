mod config;
mod routes;
mod services;
mod types;
mod utils;

use actix_web::{web::Data, App, HttpServer};
use deadpool_redis::Runtime;
use std::fs;

#[actix_web::main]
async fn main() -> Result<(), ()> {
    let _ = log4rs::init_file("log4rs.yaml", Default::default())
        .or_else::<(), _>(|err| panic!("error loading log4rs.yaml: {:?}", err));

    let config = config::load().unwrap_or_else(|err| panic!("error loading config: {:?}", err));

    let redis_pool_config = deadpool_redis::Config::from_url(config.clone().redis_url);
    let redis_pool = redis_pool_config
        .create_pool(Some(Runtime::Tokio1))
        .unwrap_or_else(|err| panic!("could not create Redis pool: {:?}", err));

    let _ = redis_pool
        .get()
        .await
        .unwrap_or_else(|err| panic!("failed to connect to Redis: {:?}", err));

    let amqp_connection = lapin::Connection::connect(
        config.amqp_url.as_str(),
        lapin::ConnectionProperties::default(),
    )
    .await
    .unwrap_or_else(|err| panic!("failed to connect to AMQP broker: {:?}", err));

    let amqp_channel = amqp_connection
        .create_channel()
        .await
        .unwrap_or_else(|err| panic!("could not create AMQP channel: {:?}", err));

    let meta_hash = fs::read_to_string(".meta/HASH")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();
    let meta_version = fs::read_to_string(".meta/VERSION")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();

    let app_context = types::AppContext {
        config: config.clone(),
        redis_pool,
        amqp_channel,
        meta: types::Meta {
            hash: meta_hash,
            version: meta_version,
        },
        services: types::Services {
            auth_service: services::AuthService::new(config.auth_service_url),
        },
    };

    let http_server = HttpServer::new(move || {
        App::new()
            .app_data(Data::new(app_context.clone()))
            .service(routes::meta_route)
            .service(routes::events_route)
            .service(routes::send_route)
            .service(routes::send_command_route)
    })
    .bind(("0.0.0.0", config.port))
    .unwrap_or_else(|err| panic!("error binding to port: {:?}", err));

    log::info!("HTTP server listening on port {}", config.port);

    let _ = http_server
        .run()
        .await
        .or_else::<(), _>(|err| panic!("error starting server: {:?}", err));

    Ok(())
}
