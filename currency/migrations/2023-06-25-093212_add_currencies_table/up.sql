CREATE TABLE IF NOT EXISTS currencies (
  id SERIAL PRIMARY KEY,
  code VARCHAR(3) NOT NULL,
  name TEXT NOT NULL,
  minor_unit_divisor INT4 NOT NULL,
  enabled BOOLEAN NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  
  CONSTRAINT currency_code_unique UNIQUE (code)
);
