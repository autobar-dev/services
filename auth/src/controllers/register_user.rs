use crate::controllers;
use crate::models;
use crate::types;
use crate::utils;

use actix_web::http::StatusCode;

pub async fn register_user_controller(
    context: types::AppContext,
    email: String,
    password: String,
    auto_login: bool,
    remember_me: bool,
    user_agent: Option<String>,
) -> Result<Option<String>, types::RestError> {
    if !utils::verify_clear_password(password.clone()) {
        return Err(types::RestError::new(
            StatusCode::BAD_REQUEST,
            "password is not sufficient",
        ));
    }

    let password_hash = utils::hash_string(password.clone());
    let register_result =
        models::UserModel::create(context.clone(), email.clone(), password_hash).await;

    if register_result.is_err() {
        return Err(types::RestError::new(
            StatusCode::BAD_REQUEST,
            "could not create user",
        ));
    }

    if !auto_login {
        return Ok(None);
    }

    let login_result = controllers::login_user_controller(
        context.clone(),
        email.clone(),
        password.clone(),
        remember_me,
        user_agent,
    )
    .await;

    if login_result.is_err() {
        return Err(types::RestError::new(
            StatusCode::INTERNAL_SERVER_ERROR,
            "could not log in after registering",
        ));
    }

    let session = login_result.unwrap();

    Ok(Some(session))
}
