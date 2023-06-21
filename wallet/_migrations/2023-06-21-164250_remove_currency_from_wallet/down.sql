-- you cannot drop an enums value after adding it

ALTER TABLE wallets
ADD COLUMN currency_code VARCHAR(3);
