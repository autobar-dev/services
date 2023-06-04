CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE transaction_type AS ENUM ('deposit', 'withdraw', 'purchase', 'refund');

CREATE TABLE IF NOT EXISTS transactions (
  id uuid DEFAULT uuid_generate_v4 (),
  wallet_id INTEGER NOT NULL,
  type transaction_type NOT NULL,
  value INTEGER NOT NULL,
  currency_code VARCHAR(3) NOT NULL,

  PRIMARY KEY (id)
);
