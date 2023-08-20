CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

  email TEXT UNIQUE NOT NULL,

  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,

  locale VARCHAR(9) NOT NULL,

  identity_verification_id TEXT DEFAULT NULL,
  identity_verification_source TEXT DEFAULT NULL,

  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

