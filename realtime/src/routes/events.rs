use actix_web::{get, web, HttpRequest, HttpResponse, Responder};
use actix_web_lab::sse;
use serde::Deserialize;

use crate::types::{
    self,
    consts::{SESSION_COOKIE_NAME, SESSION_HEADER_NAME},
    Client,
};

#[derive(Debug, Clone, Deserialize)]
pub struct EventsQuery {
    session: Option<String>,
}

#[get("/events")]
pub async fn events_route(
    req: HttpRequest,
    query: web::Query<EventsQuery>,
    data: web::Data<types::AppContext>,
) -> impl Responder {
    let context = data.as_ref().to_owned();

    log::info!("Connection info: {:?}", req.connection_info());

    let session_from_query = query.session.clone();
    let session_from_header = req.headers().get(SESSION_HEADER_NAME);
    let session_from_cookie = req.cookie(SESSION_COOKIE_NAME);

    let session: String;

    if session_from_query.is_some() {
        log::debug!("got session from query");
        session = session_from_query.unwrap();
    } else if session_from_header.is_some() {
        log::debug!("got session from header");
        session = session_from_header.unwrap().to_str().unwrap().to_string();
    } else if session_from_cookie.is_some() {
        log::debug!("got session from cookie");
        session = session_from_cookie.unwrap().value().to_string();
    } else {
        return HttpResponse::Unauthorized()
            .body("session not provided either as a cookie, query or header");
    }

    let session_data = context
        .clone()
        .services
        .auth_service
        .verify_session(session)
        .await;

    if session_data.is_err() {
        log::error!(
            "error while verifying session: {:?}",
            session_data.unwrap_err()
        );

        return HttpResponse::InternalServerError().body("could not verify session");
    }

    let session_data = session_data.unwrap();

    if session_data.is_none() {
        return HttpResponse::Unauthorized().body("session invalid");
    }

    let session_data = session_data.unwrap();

    let (sender, sse_stream) = sse::channel(2);

    let client = Client::new(
        session_data.client_type,
        session_data.client_identifier,
        sender,
    );

    let listen_result = client.clone().listen(context.clone()).await;

    if listen_result.is_err() {
        log::error!("Listen error: {:?}", listen_result.unwrap_err());

        return HttpResponse::InternalServerError().body("failed to listen");
    }

    sse_stream.respond_to(&req)
}
