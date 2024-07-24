-- +migrate Up
ALTER TABLE event_types ADD COLUMN poll_event_id TEXT;
ALTER TABLE event_types ADD COLUMN poll_contract TEXT;

-- +migrate Down
ALTER TABLE event_types DROP COLUMN poll_event_id;
ALTER TABLE event_types DROP COLUMN poll_contract;
