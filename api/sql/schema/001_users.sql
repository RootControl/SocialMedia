-- +goose Up
create extension if not exists "uuid-ossp";

CREATE TABLE users (
  id uuid PRIMARY KEY default uuid_generate_v4(),
  email text not null,
  created_at timestamp not null,
  updated_at timestamp not null
);

-- +goose Down
drop table users;
