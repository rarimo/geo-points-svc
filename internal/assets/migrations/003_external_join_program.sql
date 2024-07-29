-- +migrate Up
ALTER TABLE balances RENAME COLUMN anonymous_id TO internal_aid;
ALTER TABLE balances ADD COLUMN external_aid TEXT UNIQUE;
ALTER TABLE balances DROP COLUMN is_verified;

-- +migrate Down
ALTER TABLE balances RENAME COLUMN internal_aid TO anonymous_id;
ALTER TABLE balances DROP COLUMN external_aid;
ALTER TABLE balances ADD COLUMN is_verified BOOLEAN NOT NULL default FALSE;
UPDATE balances SET is_verified = TRUE WHERE anonymous_id IS NOT NULL;
