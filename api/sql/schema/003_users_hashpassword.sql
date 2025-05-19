-- +goose Up
alter table users add column hash_password text not null default 'unset';

-- +goose Down
alter table users drop column hash_password;
