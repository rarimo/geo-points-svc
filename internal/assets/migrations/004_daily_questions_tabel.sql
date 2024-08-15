-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_questions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    time_for_answer INTEGER NOT NULL,
    reward INTEGER NOT NULL,
    answer_options JSONB NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    starts_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

-- +migrate Down 2024-08-14 11:42:56.462615
DROP TABLE IF EXISTS daily_questions;