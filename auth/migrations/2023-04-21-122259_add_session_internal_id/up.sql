ALTER TABLE sessions
ADD COLUMN internal_id serial UNIQUE;
