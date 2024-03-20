-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS admin (
	id SERIAL PRIMARY KEY,
	uuid TEXT NOT NULL,
	login TEXT NOT NULL UNIQUE,
	password_hash TEXT
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
