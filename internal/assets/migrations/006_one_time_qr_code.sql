-- +migrate Up
CREATE OR REPLACE FUNCTION trigger_set_updated_at_ts() RETURNS trigger
    LANGUAGE plpgsql
AS $$ BEGIN NEW.updated_at = (NOW() AT TIME ZONE 'utc'); RETURN NEW; END; $$;

CREATE TABLE IF NOT EXISTS qr_codes (
    id TEXT PRIMARY KEY,
    nullifier TEXT REFERENCES balances (nullifier),
    reward BIGINT NOT NULL,
    usage_count BIGINT NOT NULL DEFAULT 0,
    infinity BOOLEAN NOT NULL DEFAULT FALSE,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TRIGGER IF EXISTS set_updated_at ON qr_codes;
CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON qr_codes
    FOR EACH ROW
EXECUTE FUNCTION trigger_set_updated_at_ts();


-- +migrate Down
DROP TABLE IF EXISTS qr_codes;
DROP FUNCTION IF EXISTS trigger_set_updated_at_ts();
