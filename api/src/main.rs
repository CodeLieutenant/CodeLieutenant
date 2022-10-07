use std::net::SocketAddr;
use std::process::exit;
use std::str::FromStr;

use ::tracing::{error, info};
use tokio::signal;

use crate::tracing::setup_tracing;

mod config;
mod handlers;
mod tracing;

#[tokio::main]
async fn main() {
    let config = config::AppConfig::new("./config").unwrap();

    let _guard = setup_tracing(&config.tokio_console);

    let addr = SocketAddr::from_str(&config.http.host).unwrap();

    let app = handlers::router();

    info!(
        addr = ?addr,
        "Axum HTTP Server is running"
    );

    let result = axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .with_graceful_shutdown(shutdown_signal())
        .await;

    if result.is_err() {
        error!(error = ?result.unwrap_err(), "Failed to create Axum web server");
    }
}

async fn shutdown_signal() {
    let ctrl_c = async {
        let result = signal::ctrl_c().await;

        match result {
            Ok(value) => {
                info!("Received Ctrl+C");
                value
            }
            Err(err) => {
                error!(error = ?err, "failed to install Ctrl+C handler");
                exit(1);
            }
        }
    };

    #[cfg(unix)]
    let terminate = async {
        let result = signal::unix::signal(signal::unix::SignalKind::terminate());

        match result {
            Ok(mut value) => {
                value.recv().await;
                info!("Received Ctrl+C");
                value
            }
            Err(err) => {
                error!(error = ?err, "failed to install Ctrl+C handler");
                exit(1);
            }
        }
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }

    info!("signal received, starting graceful shutdown");
}
