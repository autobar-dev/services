use deadpool::managed::Pool;
use deadpool_redis::{Connection, Manager};

use crate::config::Config;
use crate::services::AuthService;
use crate::types;

#[derive(Clone, Debug)]
pub struct Services {
    pub auth_service: AuthService,
}

#[derive(Clone)]
pub struct AppContext {
    pub redis_pool: Pool<Manager, Connection>,
    pub amqp_channel: lapin::Channel,
    pub config: Config,
    pub meta: types::Meta,
    pub services: Services,
}
