use envconfig::Envconfig;

#[derive(Envconfig, Debug, Clone)]
pub struct Config {
    #[envconfig(from = "DATABASE_URL")]
    pub database_url: String,

    #[envconfig(from = "PORT")]
    pub port: u16,

    #[envconfig(from = "EXCHANGE_RATE_API_KEY")]
    pub exchange_rate_api_key: String,

    #[envconfig(from = "PAST_RATE_RETENTION_SECONDS")]
    pub past_rate_retention_seconds: u32,
}

pub fn load() -> Result<Config, envconfig::Error> {
    dotenv::dotenv().ok();
    Config::init_from_env()
}

