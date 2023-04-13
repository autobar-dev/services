use crate::types;

use actix_web::{
    web,
    get,
    Responder,
    HttpResponse,
};
use serde::Serialize;

#[derive(Debug, Serialize)]
struct MetaResponse {
    status: String,
    data: types::Meta,
}

#[get("/meta")]
pub async fn meta_route(data: web::Data<types::AppContext>) -> impl Responder {
    let context = data.get_ref().clone();
    let meta = context.meta;

    HttpResponse::Ok().json(
        MetaResponse {
            status: "ok".to_string(),
            data: meta,
        }
    )
}
