CREATE TABLE IF NOT EXISTS rates (
  id SERIAL PRIMARY KEY,
  from_currency VARCHAR(3) NOT NULL,
  to_currency VARCHAR(3) NOT NULL,
  rate FLOAT8 NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  
  CONSTRAINT rate_from_to_combination_unique UNIQUE (from_currency, to_currency)
);
