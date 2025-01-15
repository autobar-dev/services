CREATE TABLE wallets (
  id SERIAL PRIMARY KEY,

  user_id INTEGER NOT NULL REFERENCES users(id),

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);