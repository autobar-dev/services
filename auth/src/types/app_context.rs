use crate::config;
use crate::types;

#[derive(Clone)]
pub struct AppContext {
    pub database_pool: sqlx::PgPool,
    pub meta_factors: types::MetaFactors,
    pub config: config::Config,
}
