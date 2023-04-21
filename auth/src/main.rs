extern crate bcrypt;

mod config;
mod controllers;
mod models;
mod types;
mod views;

use actix_web::{web, HttpServer};
use sqlx::postgres::PgPoolOptions;
use std::{fs, process};

#[actix_web::main]
async fn main() -> Result<(), ()> {
    if log4rs::init_file("log4rs.yaml", Default::default()).is_err() {
        println!("Error loading log4rs.yaml");
        process::exit(1);
    }

    let config = config::load();

    if config.is_err() {
        log::error!("Error loading config: {:?}", config.unwrap_err());
        process::exit(1);
    }

    let config = config.unwrap();

    let config_main_domain_clone = config.clone();

    if config_main_domain_clone.main_domain.is_empty() {
        log::warn!("Running without a domain specified (insecure)");
    }

    if !config_main_domain_clone.set_secure_cookies {
        log::warn!("Cookies set will not have *secure* set");
    }

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&config.database_url)
        .await;

    if pool.is_err() {
        log::error!("Error connecting to database: {:?}", pool.unwrap_err());
        process::exit(1);
    }

    let pool = pool.unwrap();

    let meta_hash = fs::read_to_string(".meta/HASH")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();
    let meta_version = fs::read_to_string(".meta/VERSION")
        .unwrap_or("".to_string())
        .trim_end()
        .to_string();

    let app_context = types::AppContext {
        database_pool: pool.clone(),
        config: config.clone(),
        meta: types::Meta {
            hash: meta_hash,
            version: meta_version,
        },
    };

    let http_server = HttpServer::new(move || {
        actix_web::App::new()
            .app_data(web::Data::new(app_context.clone()))
            .service(views::meta_route)
            .service(web::scope("/user").service(views::user::login_route))
            .service(
                web::scope("/session")
                    .service(views::session::verify_route)
                    .service(views::session::remove_route)
                    .service(views::session::all_for_user_route)
                    .service(views::session::remove_by_internal_id_route),
            )
    })
    .bind(("0.0.0.0", config.port));

    if http_server.is_err() {
        log::error!("Error binding to port: {:?}", http_server.err());
        process::exit(1);
    }

    log::info!("HTTP server listening on port {}", config.port);

    let http_server = http_server.unwrap();
    let run_result = http_server.run().await;

    if run_result.is_err() {
        log::error!("Error running server: {:?}", run_result.unwrap_err());
        process::exit(1);
    }

    Ok(())
}
