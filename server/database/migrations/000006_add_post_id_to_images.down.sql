ALTER TABLE images
DROP FOREIGN KEY fk_post_images_id;

ALTER TABLE images DROP COLUMN post_id;