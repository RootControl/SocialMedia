-- name: CreateUser :one
insert into users (id, email, created_at, updated_at, hash_password)
values (gen_random_uuid(), $1, now(), now(), $2)
returning *;

-- name: GetUserById :one
SELECT id, email, created_at, updated_at FROM users where id = $1;

-- name: GetUserByEmail :one
select id, email, hash_password, created_at, updated_at from users where email = $1;
