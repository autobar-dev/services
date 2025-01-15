CREATE TABLE currencies (
  id SERIAL PRIMARY KEY,

  code VARCHAR(3) UNIQUE NOT NULL,
  name TEXT NOT NULL,
  symbol VARCHAR(8) DEFAULT NULL,

  minor_unit_divisor INTEGER NOT NULL,
  enabled BOOLEAN NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);