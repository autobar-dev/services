ALTER TABLE modules
ADD COLUMN station_slug TEXT DEFAULT NULL,
ADD COLUMN product_slug TEXT DEFAULT NULL,
DROP COLUMN station_id,
DROP COLUMN product_id,
DROP COLUMN enabled,
DROP COLUMN updated_at;
