CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE unfinished_registrations(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

  email TEXT UNIQUE NOT NULL,
  locale VARCHAR(9) NOT NULL,
  email_confirmation_code VARCHAR(64) UNIQUE NOT NULL,
  email_confirmation_code_expires_at TIMESTAMPTZ NOT NULL,

  -- always there
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
