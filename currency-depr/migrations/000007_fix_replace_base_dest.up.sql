ALTER TABLE rates
DROP COLUMN base_currency,
DROP COLUMN destination_currency;

ALTER TABLE rates
ADD COLUMN base_currency_id INTEGER REFERENCES supported_currencies(id),
ADD COLUMN destination_currency_id INTEGER REFERENCES supported_currencies(id);