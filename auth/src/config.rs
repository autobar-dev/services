use envconfig::Envconfig;
use dotenv;

#[derive(Envconfig, Debug, Clone)]
pub struct Config {
    #[envconfig(from = "DATABASE_URL")]
    pub database_url: String,

    #[envconfig(from = "PORT")]
    pub port: u16,

    #[envconfig(from = "ALLOW_ONLY_SAME_USER_AGENT")]
    pub allow_only_same_user_agent: bool,
}

pub fn load() -> Result<Config, envconfig::Error> {
    dotenv::dotenv().ok();
    Config::init_from_env()
}
