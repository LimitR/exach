package repository

import "github.com/jmoiron/sqlx"

type Thread struct {
	db            *sqlx.DB
	Id            int64  `db:"id" json:"id"`
	Text          string `db:"text" json:"text"`
	Head          string `db:"head" json:"head"`
	Password_hash string `db:"password_hash" json:"passwordHash"`
	Img           string `db:"img" json:"img"`
}

func NewThread(db *sqlx.DB) *Thread {
	return &Thread{db: db}
}

func (t *Thread) createTableOrNotExists() {
	t.db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT,
		);`)
}

func (t *Thread) CreateThread(head, text, passwordHash, img string) (int64, error) {
	raw, _ := t.db.Exec(`
	INSERT INTO threads (head, text, password_hash, img) VALUES (?, ?, ?, ?) RETURNING id
	`, head, text, passwordHash, img)
	// if err != nil {
	// 	return "", err
	// }
	// defer raw.Close()
	// for raw.Next() {
	// 	err := raw.Scan(&id)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }
	return raw.LastInsertId()
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
