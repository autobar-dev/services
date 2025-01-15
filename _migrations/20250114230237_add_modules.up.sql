CREATE TABLE modules (
  id SERIAL PRIMARY KEY,

  serial_number VARCHAR(6) NOT NULL UNIQUE,
  private_key VARCHAR(255) NOT NULL,

  station_id INT DEFAULT NULL REFERENCES stations(id),
  product_id INT DEFAULT NULL REFERENCES products(id),

  display_unit_id INTEGER REFERENCES module_display_units(id) NOT NULL DEFAULT 1,
  display_currency_id INTEGER REFERENCES currencies(id) NOT NULL,

  prices JSON NOT NULL DEFAULT '{}', -- replace with a table
  enabled BOOLEAN NOT NULL DEFAULT true,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);