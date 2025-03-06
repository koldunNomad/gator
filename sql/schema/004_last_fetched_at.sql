-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP NULL;

-- +goose Down
ALTER Table feeds DROP COLUMN last_fetched_at;