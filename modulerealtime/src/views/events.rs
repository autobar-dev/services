use std::time::Duration;

use actix_web::{get, HttpRequest, HttpResponse, Responder};
use actix_web_lab::sse::{self, ChannelStream, Sse};

#[get("/events")]
pub async fn events_route(req: HttpRequest) -> Sse<ChannelStream> {
    log::info!("Connection info: {:?}", req.connection_info());

    let (tx, rx): (sse::Sender, Sse<ChannelStream>) = sse::channel(10);

    tokio::task::spawn(async move {
        for counter in 1..10 {
            println!("Hello #{}", counter);
            tx.send(sse::Data::new(format!("hello from server #{}", counter)))
                .await
                .unwrap();

            tokio::time::sleep(Duration::from_millis(500)).await;
        }
    });

    rx

    // HttpResponse::Ok()
    //     .insert_header(("content-type", "text/event-stream"))
    //     .streaming(rx.await)
}
