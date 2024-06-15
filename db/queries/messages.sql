-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;

-- name: CreateMessage :exec
INSERT INTO messages (
    id, 
    created_at,
    raw,
    raw_jsonb
) VALUES (
    DEFAULT, NOW(), $1, $2
);

-- name: GetMessagesAsc :many
SELECT id, created_at, raw FROM messages ORDER BY id ASC LIMIT $1 OFFSET $2;

-- name: GetMessagesDesc :many
SELECT id, created_at, raw FROM messages ORDER BY id DESC LIMIT $1 OFFSET $2;
