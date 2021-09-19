create TABLE images
(
    id         BIGSERIAL PRIMARY KEY,
    path TEXT NOT NULL,
    driver varchar(30) NOT NULL,
    link TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

ALTER TABLE images ADD COLUMN project_id BIGSERIAL;

alter table images
add CONSTRAINT fk_project_images_id
    FOREIGN KEY (project_id)
    REFERENCES projects(id);