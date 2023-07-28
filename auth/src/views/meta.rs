use crate::{types, utils};

use actix_web::{get, web, HttpResponse, Responder};
use serde::Serialize;

#[derive(Debug, Serialize)]
struct MetaResponse {
    status: String,
    data: types::Meta,
}

#[get("/meta")]
pub async fn meta_route(data: web::Data<types::AppContext>) -> impl Responder {
    let context = data.get_ref().clone();
    let meta = utils::get_meta_from_factors(context.meta_factors);

    HttpResponse::Ok().json(MetaResponse {
        status: "ok".to_string(),
        data: meta,
    })
}
