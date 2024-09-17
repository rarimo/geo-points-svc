-- +migrate Up
CREATE TABLE IF NOT EXISTS withdrawals
(
    id         uuid PRIMARY KEY default gen_random_uuid(),
    nullifier  TEXT    NOT NULL REFERENCES balances (nullifier),
    amount     integer NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS withdrawals_page_index ON withdrawals (nullifier, created_at);

-- +migrate Down
DROP TABLE IF EXISTS withdrawals;