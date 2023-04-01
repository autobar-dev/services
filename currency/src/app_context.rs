use crate::types;

#[derive(Clone)]
pub struct Context {
  pub database_pool: sqlx::PgPool,
  pub meta: types::Meta,
}