CREATE TABLE display_units (
  id SERIAL PRIMARY KEY,
  symbol TEXT NOT NULL,
  divisor_from_millilitres REAL NOT NULL,
  decimals_displayed INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

-- add litre as default display_unit
INSERT INTO display_units (
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

ALTER TABLE modules
ADD COLUMN display_unit_id INTEGER REFERENCES display_units(id) NOT NULL DEFAULT 1;
