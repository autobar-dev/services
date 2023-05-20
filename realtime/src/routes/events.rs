use actix_web::{get, web, HttpRequest, HttpResponse, Responder};
use actix_web_lab::sse;

use crate::types::{self, Client, ClientType};

#[get("/events/{id}")]
pub async fn events_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
    path: web::Path<String>,
) -> impl Responder {
    let context = data.as_ref().to_owned();
    let id = path.into_inner();

    log::info!("Connection info: {:?}", req.connection_info());

    let (sender, sse_stream) = sse::channel(2);

    let client = Client::new(ClientType::Module, id.clone(), sender);

    let listen_result = client.clone().listen(context.clone()).await;

    if listen_result.is_err() {
        log::error!("Listen error: {:?}", listen_result.unwrap_err());

        return HttpResponse::InternalServerError().body("failed to listen");
    }

    sse_stream.respond_to(&req)
}
