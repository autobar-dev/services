ALTER TABLE wallets
DROP COLUMN currency_code;

ALTER TYPE transaction_type ADD VALUE IF NOT EXISTS 'currency_change';
