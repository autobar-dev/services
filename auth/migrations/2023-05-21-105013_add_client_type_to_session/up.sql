CREATE TYPE client_type
AS ENUM ('user', 'module');

ALTER TABLE sessions
ADD COLUMN client_type client_type NOT NULL DEFAULT 'user';
