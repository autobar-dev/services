use crate::app_context;
use crate::models;
use crate::types::RestError;

pub async fn new_currency_controller(
    context: app_context::Context,
    code: String,
    name: String,
    minor_unit_divisor: i32,
    symbol: Option<String>,
    enabled: bool,
) -> Result<(), RestError> {
    let context = context.clone();
    let currency =
        models::CurrencyModel::create(context, code, name, minor_unit_divisor, symbol, enabled)
            .await;

    if currency.is_err() {
        return Err(RestError::new("Error creating new currency".to_string()));
    }

    if currency.unwrap() == 0 {
        return Err(RestError::new("Error creating new currency".to_string()));
    }

    Ok(())
}
