CREATE TABLE modules (
  id SERIAL PRIMARY KEY,

  serial_number VARCHAR(6) NOT NULL UNIQUE,
  private_key VARCHAR(255) NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
