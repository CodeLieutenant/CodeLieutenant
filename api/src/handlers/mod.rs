use axum::Router;

mod contact;

#[inline]
pub(crate) fn router() -> Router {
    Router::new()
        .nest("/contact", contact::router())
}
