
alter table images
drop CONSTRAINT fk_project_images_id;

DROP TABLE IF EXISTS images;