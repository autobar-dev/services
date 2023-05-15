use actix_web::{post, HttpRequest, HttpResponse, Responder};

// #[]
// pub struct ActivateBody {
//
// }

#[post("/activate")]
pub async fn activate_route(req: HttpRequest) -> impl Responder {
    HttpResponse::Ok()
}
