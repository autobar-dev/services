use crate::app_context;
use crate::models;
use crate::types;
use crate::types::error;

pub async fn get_remote_rate_controller(
    context: app_context::Context,
    from: String,
    to: String,
) -> Result<types::RemoteRate, error::RestError> {
    let context = context.clone();
    let rate = models::RemoteRateModel::get(context, from, to).await;

    if rate.is_err() {
        let rate_error = rate.unwrap_err();

        return Err(error::RestError::new(rate_error.message));
    }

    let rate = types::RemoteRate::from(rate.unwrap());

    Ok(rate)
}

