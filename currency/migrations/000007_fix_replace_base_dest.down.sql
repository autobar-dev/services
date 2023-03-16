ALTER TABLE rates
DROP COLUMN base_currency_id,
DROP COLUMN destination_currency_id;

ALTER TABLE rates
ADD COLUMN base_currency CHAR(3) NOT NULL,
ADD COLUMN destination_currency CHAR(3) NOT NULL;