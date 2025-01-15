CREATE TABLE files (
  id SERIAL PRIMARY KEY,

  s3_object_id TEXT NOT NULL,

  name TEXT NOT NULL,
  checksum TEXT NOT NULL,
  size INTEGER NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);