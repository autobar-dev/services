CREATE TABLE unfinished_registrations (
  id SERIAL PRIMARY KEY,

  email TEXT NOT NULL UNIQUE,
  locale_id INTEGER NOT NULL REFERENCES locales(id),

  email_confirmation_code VARCHAR(64) NOT NULL UNIQUE,
  email_confirmation_code_expires_at TIMESTAMPTZ NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);