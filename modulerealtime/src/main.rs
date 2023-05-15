mod config;
mod routes;
mod types;

use actix_web::{web::Data, App, HttpServer};
use deadpool_redis::Runtime;
use lapin;

#[actix_web::main]
async fn main() -> Result<(), ()> {
    let _ = log4rs::init_file("log4rs.yaml", Default::default())
        .or_else::<(), _>(|err| panic!("error loading log4rs.yaml: {:?}", err));

    let config = config::load().unwrap_or_else(|err| panic!("error loading config: {:?}", err));

    let redis_pool_config = deadpool_redis::Config::from_url(config.redis_url);
    let redis_pool = redis_pool_config
        .create_pool(Some(Runtime::Tokio1))
        .unwrap_or_else(|err| panic!("could not create Redis pool: {:?}", err));

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

    let app_context = types::AppContext {
        redis_pool,
        amqp_channel,
    };

    let http_server = HttpServer::new(move || {
        App::new()
            .app_data(Data::new(app_context.clone()))
            .service(routes::events_route)
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
