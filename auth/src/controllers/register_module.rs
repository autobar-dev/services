use crate::models;
use crate::models::ModuleModel;
use crate::types;
use crate::types::consts::PRIVATE_KEY_LENGTH;
use crate::types::CreateModuleResponse;
use crate::utils;

use actix_web::http::StatusCode;

pub async fn register_module_controller(
    context: types::AppContext,
    serial_number: String,
) -> Result<CreateModuleResponse, types::RestError> {
    let private_key = utils::generate_private_key(PRIVATE_KEY_LENGTH);

    let private_key_hash = utils::hash_string(private_key.clone());
    let register_result =
        models::ModuleModel::create(context.clone(), serial_number.clone(), private_key_hash).await;

    if register_result.is_err() {
        return Err(types::RestError::new(
            StatusCode::BAD_REQUEST,
            "could not create module",
        ));
    }

    let module_result = ModuleModel::get_by_serial_number(context.clone(), serial_number).await;
    if module_result.is_err() {
        return Err(types::RestError::new(
            StatusCode::BAD_REQUEST,
            "could not get created module",
        ));
    }

    let create_module_response =
        CreateModuleResponse::from(module_result.unwrap(), private_key.clone());

    Ok(create_module_response)
}
