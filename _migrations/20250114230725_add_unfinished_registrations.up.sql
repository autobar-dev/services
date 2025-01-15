CREATE TABLE unfinished_registrations (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  email TEXT NOT NULL UNIQUE,
  locale VARCHAR(9) NOT NULL,
  email_confirmation_code VARCHAR(64) NOT NULL UNIQUE,
  email_confirmation_code_expires_at TIMESTAMPTZ NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);