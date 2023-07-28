extern crate chrono;

mod app_context;
mod config;
mod controllers;
mod models;
mod types;
mod utils;
mod views;

use actix_web::{web, HttpServer};

use sqlx::{
    postgres::{PgConnectOptions, PgPoolOptions},
    ConnectOptions,
};

use std::{process, str::FromStr};

#[actix_web::main]
async fn main() -> Result<(), ()> {
    // Logger
    if log4rs::init_file("log4rs.yaml", Default::default()).is_err() {
        println!("Error loading log4rs.yaml");
        process::exit(1);
    }

    // Config
    let config = config::load();

    if config.is_err() {
        log::error!("Error loading config: {:?}", config.err());
        process::exit(1);
    }

    let config = config.unwrap();

    // Database connection
    let pg_options = PgConnectOptions::from_str(&config.database_url)
        .unwrap_or_else(|err| panic!("failed to parse database URI: {}", err))
        .disable_statement_logging()
        .clone();

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect_with(pg_options)
        .await;

    if pool.is_err() {
        log::error!("Error connecting to database: {:?}", pool.err());
        process::exit(1);
    }

    let pool = pool.unwrap();

    // Create context
    let app_context = app_context::Context {
        database_pool: pool.clone(),
        config: config.clone(),
        meta_factors: utils::get_meta_factors(),
    };

    // Server
    let http_server = HttpServer::new(move || {
        actix_web::App::new()
            .app_data(web::Data::new(app_context.clone()))
            .service(views::meta_route)
            .service(
                web::scope("/currency")
                    .service(views::currency::index_route)
                    .service(views::currency::new_route)
                    .service(views::currency::enabled_route)
                    .service(views::currency::all_route)
                    .service(views::currency::set_enabled_route)
                    .service(views::currency::delete_route),
            )
            .service(
                web::scope("/rate")
                    .service(views::rate::index_route)
                    .service(views::rate::set_route)
                    .service(views::rate::remote_route)
                    .service(views::rate::delete_route)
                    .service(views::rate::update_route),
            )
    })
    .bind(("0.0.0.0", config.port));

    if http_server.is_err() {
        log::error!("Error binding to port: {:?}", http_server.err());
        process::exit(1);
    }

    log::info!("HTTP Server listening on port {}", config.port);

    let http_server = http_server.unwrap();
    let run_result = http_server.run().await;

    if run_result.is_err() {
        log::error!("Error running server: {:?}", run_result.err());
        process::exit(1);
    }

    Ok(())
}
