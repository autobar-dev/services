CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email text NOT NULL ,
  password text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT current_timestamp,

  UNIQUE (email)
);

