use std::sync::Arc;

use config::{Config, ConfigError, File, FileFormat};
use serde_derive::Deserialize;

#[derive(Debug, Deserialize)]
pub(crate) struct Http {
    pub(crate) host: String,
}

#[derive(Debug, Deserialize)]
pub(crate) struct TokioConsole {
    pub(crate) bind: String,
    pub(crate) publish_interval_sec: u64,
    pub(crate) retention_sec: u64,
}

#[derive(Debug, Deserialize)]
pub(crate) struct AppConfig {
    pub(crate) http: Http,
    pub(crate) tokio_console: TokioConsole,
}

impl AppConfig {
    pub(crate) fn new<'a>(config_path: impl Into<&'a str>) -> Result<Arc<Self>, ConfigError> {
        let file_source = File::with_name(config_path.into())
            .required(true)
            .format(FileFormat::Toml);

        let env_source = config::Environment::with_prefix("MALUSEV_API")
            .try_parsing(true)
            .separator("_")
            .list_separator(",")
            .ignore_empty(false)
            .prefix_separator("__")
            .keep_prefix(false);

        let config = Config::builder()
            .add_source(file_source)
            .add_source(env_source)
            .build()?;

        let config: Self = config.try_deserialize()?;

        Ok(Arc::new(config))
    }
}
