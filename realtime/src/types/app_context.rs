use deadpool::managed::Pool;
use deadpool_redis::{Connection, Manager};
use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use crate::config::Config;

use super::Client;

#[derive(Clone)]
pub struct AppContext {
    pub redis_pool: Pool<Manager, Connection>,
    pub amqp_channel: lapin::Channel,
    pub config: Config,
}
