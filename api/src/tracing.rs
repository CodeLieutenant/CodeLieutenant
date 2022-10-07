use std::net::SocketAddr;
use std::str::FromStr;
use std::time::Duration;

use tracing_subscriber::{filter::EnvFilter, fmt::layer as fmt_layer, prelude::*, registry};

use console_subscriber::ConsoleLayer;
use tracing_appender::non_blocking as nb_appender;
use tracing_appender::non_blocking::WorkerGuard;

pub(crate) struct Guard {
    _stdout_guard: WorkerGuard,
}

pub(crate) fn setup_tracing(config: &crate::config::TokioConsole) -> Guard {
    let env_filter = EnvFilter::try_from_default_env()
        .or_else(|_| EnvFilter::try_new("info"))
        .unwrap();
    let (stdout_non_blocking, stdout_guard) = nb_appender(std::io::stdout());

    let stdout_layer = fmt_layer()
        .with_ansi(true)
        .with_level(true)
        .with_thread_names(false)
        .with_target(true)
        .with_line_number(true)
        .with_writer(stdout_non_blocking);

    #[cfg(debug_assertions)]
    let stdout_layer = stdout_layer.pretty();
    #[cfg(not(debug_assertions))]
    let stdout_layer = stdout_layer.json();

    let layer = ConsoleLayer::builder()
        .enable_self_trace(false)
        .server_addr(SocketAddr::from_str(&config.bind).unwrap())
        .publish_interval(Duration::from_secs(config.publish_interval_sec))
        .retention(Duration::from_secs(config.retention_sec))
        .spawn();

    registry()
        .with(env_filter)
        .with(stdout_layer)
        .with(layer)
        .init();

    Guard {
        _stdout_guard: stdout_guard,
    }
}
