CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id SERIAL PRIMARY KEY,

  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,

  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,

  role VARCHAR(32) NOT NULL DEFAULT 'user',

  locale VARCHAR(9) NOT NULL,

  identity_verification_id TEXT DEFAULT NULL,
  identity_verification_source TEXT DEFAULT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
