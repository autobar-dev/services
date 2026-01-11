CREATE TABLE users (
  id SERIAL PRIMARY KEY,

  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,

  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  locale_id INTEGER NOT NULL DEFAULT 1 REFERENCES locales(id),

  role_id INTEGER NOT NULL DEFAULT 1 REFERENCES user_roles(id),

  identity_verification_id INTEGER DEFAULT NULL REFERENCES user_identity_verifications(id),

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
