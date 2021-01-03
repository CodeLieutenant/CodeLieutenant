CREATE TABLE contacts
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT,
    subject    TEXT,
    email      TEXT,
    message    TEXT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE contacts
    ADD COLUMN search tsvector;


UPDATE contacts
SET search = to_tsvector('english', subject || ' ' || message);


CREATE
INDEX contact_search_idx ON contacts USING GIN(search);

