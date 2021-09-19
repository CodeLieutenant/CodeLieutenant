CREATE TABLE IF NOT EXISTS posts
(
    id         BIGSERIAL PRIMARY KEY,
    title varchar(260) NOT NULL,
    slug varchar(300) NOT NULL,
    content TEXT NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE UNIQUE INDEX unique_title_idx ON posts (title);
CREATE UNIQUE INDEX unique_slug_idx ON posts (slug);