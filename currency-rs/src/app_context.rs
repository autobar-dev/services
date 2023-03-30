#[derive(Clone)]
pub struct Context {
  pub database_pool: sqlx::PgPool,
  pub message: String,
}