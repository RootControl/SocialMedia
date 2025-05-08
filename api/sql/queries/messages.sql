-- name: SaveMessage :one
INSERT INTO messages (id, body, user_id, created_at, updated_at)
VALUES (gen_random_uuid(), $1, $2, now(), now())
returning *;

-- name: GetMessages :many
select id, body, user_id, created_at, updated_at from messages;

-- name: GetMessage :one
select id, body, user_id, created_at, updated_at from messages where id = $1;

