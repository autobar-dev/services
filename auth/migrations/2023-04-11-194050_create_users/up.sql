CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id int4 NOT NULL,
  user_agent text DEFAULT NULL,
  valid_until timestamptz NOT NULL,
  last_used timestamptz NOT NULL DEFAULT current_timestamp,
  created_at timestamptz NOT NULL DEFAULT current_timestamp
);
