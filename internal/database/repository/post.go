package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Post struct {
	db            *sqlx.DB
	Id            int64          `db:"id"`
	Text          string         `db:"text"`
	Thread_id     sql.NullString `db:"thread_id"`
	Password_hash sql.NullString `db:"password_hash"`
	Img           sql.NullString `db:"img"`
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{
		db: db,
	}
}

func (t *Post) createTableOrNotExists() error {
	_, err := t.db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT
		);`)
	return err
}

func (t *Post) CreatePost(text, img, thread_id string) error {
	_, err := t.db.Exec(`
	INSERT INTO threads (text, img, thread_id) VALUES (?, ?, ?)
	`, text, img, thread_id)
	return err
}
