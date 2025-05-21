-- name: SaveRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expired_at, revoked_at)
VALUES ($1, now(), now(), $2, $3, $4)
returning *;

-- name: GetRefreshToken :one
select token, created_at, updated_at, user_id, expired_at, revoked_at from refresh_tokens where user_id = $1;
