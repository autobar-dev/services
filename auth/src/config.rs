use envconfig::Envconfig;


#[derive(Envconfig, Debug, Clone)]
pub struct Config {
    #[envconfig(from = "DATABASE_URL")]
    pub database_url: String,

    #[envconfig(from = "PORT")]
    pub port: u16,

    #[envconfig(from = "ALLOW_ONLY_SAME_USER_AGENT")]
    pub allow_only_same_user_agent: bool,

    #[envconfig(from = "REMEMBER_ME_DURATION_SECONDS")]
    pub remember_me_duration_seconds: i64,

    #[envconfig(from = "DEFAULT_SESSION_DURATION_SECONDS")]
    pub default_session_duration_seconds: i64,

    #[envconfig(from = "MAIN_DOMAIN")]
    pub main_domain: String,

    #[envconfig(from = "SET_SECURE_COOKIES")]
    pub set_secure_cookies: bool,
}

pub fn load() -> Result<Config, envconfig::Error> {
    dotenv::dotenv().ok();
    Config::init_from_env()
}
