use ::tracing::info;

use crate::tracing::setup_tracing;

mod tracing;

fn main() {
    let _guard = setup_tracing();

    info!(name = "Dusan", "Hello, world!");
}
