-- name: CreateUser :one
insert into users (id, email, created_at, updated_at)
values (gen_random_uuid(), $1, now(), now())
returning *;
