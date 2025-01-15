CREATE TABLE wallet_transaction_types (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

INSERT INTO wallet_transaction_types (id, name) VALUES (1, 'deposit');
INSERT INTO wallet_transaction_types (id, name) VALUES (2, 'withdraw');
INSERT INTO wallet_transaction_types (id, name) VALUES (3, 'purchase');
INSERT INTO wallet_transaction_types (id, name) VALUES (4, 'refund');
INSERT INTO wallet_transaction_types (id, name) VALUES (5, 'currency_change');

CREATE TABLE wallet_transactions (
  id SERIAL PRIMARY KEY,

  wallet_id INTEGER NOT NULL REFERENCES wallets(id),
  
  type_id INTEGER NOT NULL REFERENCES wallet_transaction_types(id),
  value INTEGER NOT NULL,
  currency_id INTEGER NOT NULL REFERENCES currencies(id),

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);