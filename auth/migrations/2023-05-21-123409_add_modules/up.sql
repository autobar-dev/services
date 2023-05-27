CREATE TABLE IF NOT EXISTS modules (
  id serial PRIMARY KEY,
  serial_number VARCHAR(6) NOT NULL,
  private_key text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT current_timestamp,

  UNIQUE (serial_number)
);

