use crate::app_context;
use crate::models;
use crate::types::RestError;

pub async fn delete_rate_controller(
    context: app_context::Context,
    from: String,
    to: String,
) -> Result<(), RestError> {
    let context = context.clone();

    let delete_result = models::RateModel::delete(context, from, to).await;

    if delete_result.is_err() {
        return Err(RestError::new("Error deleting rate".to_string()));
    }

    if delete_result.unwrap() == 0 {
        return Err(RestError::new("Rate not found".to_string()));
    }

    Ok(())
}
