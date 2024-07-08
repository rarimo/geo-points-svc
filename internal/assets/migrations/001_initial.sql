-- +migrate Up
CREATE OR REPLACE FUNCTION trigger_set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
AS $$ BEGIN NEW.updated_at = EXTRACT('EPOCH' FROM NOW()); RETURN NEW; END; $$;

CREATE TABLE IF NOT EXISTS balances
(
    nullifier          TEXT PRIMARY KEY,
    amount             bigint  NOT NULL default 0,
    created_at         integer NOT NULL default EXTRACT('EPOCH' FROM NOW()),
    updated_at         integer NOT NULL default EXTRACT('EPOCH' FROM NOW()),
    referred_by        text,
    level              INT     NOT NULL,
    anonymous_id       TEXT UNIQUE,
    shared_hash        TEXT UNIQUE,
    is_verified        BOOLEAN NOT NULL default FALSE
);

CREATE INDEX IF NOT EXISTS balances_page_index ON balances (amount, updated_at) WHERE referred_by IS NOT NULL;

DROP TRIGGER IF EXISTS set_updated_at ON balances;
CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON balances
    FOR EACH ROW
EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE IF NOT EXISTS referrals
(
    id         text PRIMARY KEY,
    nullifier  TEXT    NOT NULL REFERENCES balances (nullifier),
    usage_left INTEGER NOT NULL DEFAULT 1,
    infinity   BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS referrals_nullifier_index ON referrals (nullifier);

CREATE TABLE IF NOT EXISTS events
(
    id            uuid PRIMARY KEY NOT NULL default gen_random_uuid(),
    nullifier     TEXT             NOT NULL REFERENCES balances (nullifier),
    type          text             NOT NULL,
    status        text             NOT NULL,
    created_at    integer          NOT NULL default EXTRACT('EPOCH' FROM NOW()),
    updated_at    integer          NOT NULL default EXTRACT('EPOCH' FROM NOW()),
    meta          jsonb,
    points_amount integer,
    external_id   text,
    CONSTRAINT unique_external_id UNIQUE (nullifier, type, external_id)
);

CREATE INDEX IF NOT EXISTS events_page_index ON events (nullifier, updated_at);

DROP TRIGGER IF EXISTS set_updated_at ON events;
CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON events
    FOR EACH ROW
EXECUTE FUNCTION trigger_set_updated_at();

CREATE TABLE IF NOT EXISTS event_types
(
    name              text PRIMARY KEY NOT NULL,
    short_description text             NOT NULL,
    description       text             NOT NULL,
    reward            integer          NOT NULL,
    title             text             NOT NULL,
    frequency         text             NOT NULL,
    starts_at         timestamp,
    expires_at        timestamp,
    no_auto_open      BOOLEAN          NOT NULL DEFAULT FALSE,
    auto_claim        BOOLEAN          NOT NULL DEFAULT FALSE,
    disabled          BOOLEAN          NOT NULL DEFAULT FALSE,
    action_url        text,
    logo              text,
    qr_code_value     text
);

-- +migrate Down
DROP TABLE IF EXISTS event_types;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS referrals;
DROP TABLE IF EXISTS balances;

DROP FUNCTION IF EXISTS trigger_set_updated_at();
