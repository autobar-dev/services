use deadpool::managed::Pool;
use deadpool_redis::{Connection, Manager};

use crate::config::Config;

#[derive(Clone)]
pub struct AppContext {
    pub redis_pool: Pool<Manager, Connection>,
    pub amqp_channel: lapin::Channel,
    pub config: Config,
}
