CREATE TABLE products (
  id SERIAL PRIMARY KEY,

  slug TEXT NOT NULL UNIQUE,
  image_file_id INTEGER DEFAULT NULL REFERENCES files(id),
  enabled BOOLEAN NOT NULL DEFAULT false,

  names JSON NOT NULL DEFAULT '{}', -- replace with table
  descriptions JSON NOT NULL DEFAULT '{}', -- replace with table
  badges JSON NOT NULL DEFAULT '[]', -- replace with table

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);