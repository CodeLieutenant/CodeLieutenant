CREATE TABLE contacts(id BIGSERIAL PRIMARY KEY,
                                           name TEXT, surname TEXT, subject TEXT, email TEXT, message TEXT);


ALTER TABLE contacts ADD COLUMN search tsvector;


UPDATE contacts
SET search = to_tsvector('english', subject || ' ' || message);


CREATE INDEX contact_search_idx ON contacts USING GIN(search);

