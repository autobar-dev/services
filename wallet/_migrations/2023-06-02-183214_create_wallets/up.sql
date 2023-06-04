CREATE TABLE IF NOT EXISTS wallets (
  id serial PRIMARY KEY,
  user_email TEXT UNIQUE NOT NULL,
  currency_code VARCHAR(3) NOT NULL
);
