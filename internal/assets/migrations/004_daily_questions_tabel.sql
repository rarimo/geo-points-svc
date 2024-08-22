-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_questions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    time_for_answer INTEGER NOT NULL,
    reward INTEGER NOT NULL,
    answer_options JSONB NOT NULL,
    starts_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    correct_answer INTEGER NOT NULL,
    num_correct_answers INTEGER DEFAULT 0,
    num_incorrect_answers INTEGER DEFAULT 0,
    num_no_answer INTEGER DEFAULT 0,
    num_all_participants INTEGER DEFAULT 0
);

-- +migrate Down
DROP TABLE IF EXISTS daily_questions;