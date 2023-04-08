use crate::app_context;
use crate::controllers;
use crate::types::error;

pub async fn update_rate_controller(
    context: app_context::Context,
    from: String,
    to: String,
) -> Result<(), error::RestError> {
    let context = context.clone();

    let context_remote_rate_clone = context.clone();
    let from_remote_rate_clone = from.clone();
    let to_remote_rate_clone = to.clone();

    let remote_rate = controllers::get_remote_rate_controller(
        context_remote_rate_clone,
        from_remote_rate_clone,
        to_remote_rate_clone,
    )
    .await;

    if remote_rate.is_err() {
        let remote_rate_error = remote_rate.unwrap_err();

        log::error!("Error updating rate: {:?}", remote_rate_error);
        return Err(error::RestError::new(remote_rate_error.message));
    }

    let remote_rate = remote_rate.unwrap();

    let set_rate = controllers::set_rate_controller(context, from, to, remote_rate.rate).await;

    if set_rate.is_err() {
        let set_rate_error = set_rate.unwrap_err();

        log::error!("Error updating rate: {:?}", set_rate_error);
        return Err(set_rate_error);
    }

    Ok(())
}
