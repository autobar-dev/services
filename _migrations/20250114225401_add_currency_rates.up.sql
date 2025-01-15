CREATE TABLE currency_rates (
  id SERIAL PRIMARY KEY,

  from_currency_id INTEGER NOT NULL REFERENCES currencies(id),
  to_currency_id INTEGER NOT NULL REFERENCES currencies(id),
  rate FLOAT8 NOT NULL,

  fetched_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);