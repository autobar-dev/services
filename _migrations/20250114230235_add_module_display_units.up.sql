CREATE TABLE module_display_units (
  id SERIAL PRIMARY KEY,

  symbol TEXT NOT NULL,
  amount REAL NOT NULL DEFAULT 1,
  divisor_from_millilitres REAL NOT NULL,
  decimals_displayed INTEGER NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

-- add litre as default display_unit
INSERT INTO module_display_units (
  id,
  symbol,
  divisor_from_millilitres,
  decimals_displayed
) VALUES (
  1,
  'L',
  1000,
  3
);