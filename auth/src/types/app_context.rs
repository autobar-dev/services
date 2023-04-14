use crate::types;
use crate::config;

#[derive(Clone)]
pub struct AppContext {
    pub database_pool: sqlx::PgPool,
    pub meta: types::Meta,
    pub config: config::Config,
}
