CREATE TABLE user_identity_verifications (
  id SERIAL PRIMARY KEY,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);