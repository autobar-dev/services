CREATE TABLE IF NOT EXISTS products (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

  names JSON NOT NULL DEFAULT '{}',
  descriptions JSON NOT NULL DEFAULT '{}',
  cover TEXT DEFAULT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
