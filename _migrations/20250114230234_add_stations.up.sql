CREATE TABLE stations (
  id SERIAL PRIMARY KEY,

  slug TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);