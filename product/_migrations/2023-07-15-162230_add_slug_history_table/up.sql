CREATE TABLE IF NOT EXISTS slug_history (
  id SERIAL PRIMARY KEY,

  product_id uuid NOT NULL,
  slug VARCHAR(64) NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

  CONSTRAINT fk_product_id
    FOREIGN KEY(product_id)
      REFERENCES products(id)
);
