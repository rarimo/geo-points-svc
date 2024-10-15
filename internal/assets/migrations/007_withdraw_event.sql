-- +migrate Up
INSERT INTO events(nullifier, type, status, created_at, updated_at, points_amount)
SELECT w.nullifier AS nullifier, 'withdraw' AS type, 'claimed' AS status, 
EXTRACT('EPOCH' FROM w.created_at) AS created_at, EXTRACT('EPOCH' FROM w.created_at) AS updated_at,
(-w.amount) AS points_amount FROM withdrawals w;

-- +migrate Down
DELETE FROM events WHERE type='withdraw';
