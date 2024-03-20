-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	head TEXT,
	text TEXT NOT NULL,
	thread_id TEXT,
	password_hash TEXT,
	img TEXT
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
