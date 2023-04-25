mod config;
mod views;

use actix_web::{App, HttpServer};

#[actix_web::main]
async fn main() -> Result<(), ()> {
    let _ = log4rs::init_file("log4rs.yaml", Default::default())
        .or_else::<(), _>(|err| panic!("error loading log4rs.yaml: {:?}", err));

    let config = config::load().unwrap_or_else(|err| panic!("error loading config: {:?}", err));

    let http_server = HttpServer::new(move || App::new().service(views::events_route))
        .bind(("0.0.0.0", config.port))
        .unwrap_or_else(|err| panic!("error binding to port: {:?}", err));

    log::info!("HTTP server listening on port {}", config.port);

    let _ = http_server
        .run()
        .await
        .or_else::<(), _>(|err| panic!("error starting server: {:?}", err));

    Ok(())
}
