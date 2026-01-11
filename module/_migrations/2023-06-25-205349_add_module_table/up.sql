CREATE TABLE IF NOT EXISTS modules (
  id SERIAL PRIMARY KEY,
  serial_number VARCHAR(6) UNIQUE NOT NULL,
  
  station_slug TEXT DEFAULT NULL,
  product_slug TEXT DEFAULT NULL,

  prices JSON NOT NULL DEFAULT '{}', -- replace with a table
  
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
