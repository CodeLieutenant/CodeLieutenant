CREATE TABLE subscriptions
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT,
    email      TEXT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE UNIQUE INDEX unique_email_idx ON subscriptions (email);