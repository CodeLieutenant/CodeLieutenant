create TABLE projects
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    link TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);