use actix_web::http::StatusCode;

use crate::models;
use crate::types;

pub async fn _get_user_controller(
    context: types::AppContext,
    email: String,
) -> Result<types::User, types::RestError> {
    let user = models::UserModel::get_by_email(context, email).await;

    if user.is_err() {
        return Err(types::RestError::new(
            StatusCode::NOT_FOUND,
            "user not found",
        ));
    }

    let user = user.unwrap();

    Ok(types::User::from(user))
}
