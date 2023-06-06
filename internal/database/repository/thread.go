package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Thread struct {
	db            *sqlx.DB
	ctx           context.Context
	Id            int64  `db:"id" json:"id"`
	Text          string `db:"text" json:"text"`
	Head          string `db:"head" json:"head"`
	Password_hash string `db:"password_hash" json:"passwordHash"`
	Img           string `db:"img" json:"img"`
}

func NewThread(db *sqlx.DB, ctx context.Context) *Thread {
	newCtx, _ := context.WithCancel(ctx)
	return &Thread{db: db, ctx: newCtx}
}

func (t *Thread) createTableOrNotExists() {
	t.db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id SERIAL PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT
		);`)
}

func (t *Thread) CreateThread(head, text, passwordHash, img string) (int, error) {
	lastInsertId := 0
	err := t.db.QueryRow(`
	INSERT INTO threads (head, text, password_hash, img) VALUES ($1, $2, $3, $4) RETURNING id
	`, head, text, passwordHash, img).Scan(&lastInsertId)
	return lastInsertId, err
}

func (t *Thread) CheckPassword(id, passwordHash string) bool {
	result := Thread{}
	t.db.Get(&result, `SELECT password_hash FROM threads WHERE id = $1`, id)
	return result.Password_hash == passwordHash
}

func (t *Thread) GetThreads(limit int32) []Thread {
	result := []Thread{}
	t.db.Select(
		&result,
		`SELECT id, head, text, img FROM threads WHERE thread_id IS NULL LIMIT $1`,
		limit,
	)
	return result
}

func (t *Thread) GetThreadAndPosts(threadId string, limit int32) []Thread {
	result := []Thread{}
	t.db.Select(
		&result,
		`SELECT id, head, text, img FROM threads WHERE id = $1 OR thread_id = $1 LIMIT $2`,
		threadId,
		limit,
	)
	return result
}
