-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_questions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    time_for_answer INTEGER NOT NULL,
    bounty INTEGER NOT NULL,
    answer_options JSONB NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    starts_at INTEGER
);

-- +migrate Down
DROP TABLE IF EXISTS daily_questions;
