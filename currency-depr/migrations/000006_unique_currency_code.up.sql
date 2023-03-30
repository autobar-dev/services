ALTER TABLE supported_currencies
ADD CONSTRAINT unique_currency_code UNIQUE (code);