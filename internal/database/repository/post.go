package repository

import "database/sql"

type Post struct {
	text          string `db:"text"`
	thread_id     string `db:"thread_id"`
	password_hash string `db:"password_hash"`
	img           string `db:"img"`
}

func (t *Post) createTableOrNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT
		);`)
	return err
}

func (t *Post) CreatePost(db *sql.DB, text, img, thread_id string) error {
	_, err := db.Exec(`
	INSERT INTO threads (text, img, thread_id) VALUES ($1, $2, $3)
	`, text, img, thread_id)
	return err
}
