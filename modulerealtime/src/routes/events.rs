use std::time::Duration;

use actix_web::{get, web, HttpRequest};
use actix_web_lab::sse::{self, ChannelStream, Sse};

use crate::types::{self, ModuleConnection};

#[get("/events")]
pub async fn events_route(
    req: HttpRequest,
    data: web::Data<types::AppContext>,
) -> Sse<ChannelStream> {
    let context = data.as_ref().to_owned();
    log::info!("Connection info: {:?}", req.connection_info());

    let (sender, sse_stream) = sse::channel(1);

    let serial_number = "1234567890ABCDEF".to_string();

    let mc = ModuleConnection::new(serial_number, sender).await;

    for counter in 1..10 {
        let _ = mc
            .sse_sender
            .send(sse::Event::Comment(format!("hello {}", counter).into()))
            .await;
    }

    let _ = mc.on_disconnect(context.clone()).await;

    sse_stream
}
