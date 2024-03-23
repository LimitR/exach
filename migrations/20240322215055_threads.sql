-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS threads (
	id SERIAL PRIMARY KEY,
	head TEXT DEFAULT '',
	text TEXT NOT NULL,
	user_name TEXT DEFAULT 'Anonymous',
	img TEXT DEFAULT '',
    thread_id TEXT default ''
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
