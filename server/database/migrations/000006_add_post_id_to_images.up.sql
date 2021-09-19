ALTER TABLE images ADD COLUMN post_id BIGSERIAL;

alter table images
add CONSTRAINT fk_post_images_id
    FOREIGN KEY (post_id)
    REFERENCES posts(id);