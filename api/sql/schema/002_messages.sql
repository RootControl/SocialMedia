-- +goose Up
CREATE TABLE messages (
  id uuid PRIMARY KEY default uuid_generate_v4(),
  body text not null,
  user_id uuid not null references users (id) on delete cascade,
  created_at timestamp not null,
  updated_at timestamp not null
);


-- +goose Down
DROP TABLE messages;
