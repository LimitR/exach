package threads

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

type Thread struct {
	Db *sql.DB
}

func NewThread() *Thread {
	secret, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open(secret["DB_DRIVER"], secret["DATA_SOURCE_NAME"])
	if err != nil {
		log.Fatalln(err)
	}
	t := &Thread{
		Db: db,
	}
	t.createTableOrNotExists()
	return t
}

func (t *Thread) createTableOrNotExists() {
	t.Db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT,
		);`)
}

func (t *Thread) CreateThread(head, text, passwordHash, img string) (string, error) {
	var id string
	raw, err := t.Db.Query(`
	INSERT INTO threads (head, text, password_hash, img) VALUES ($1, $2, $3, $4) RETURNING id
	`, head, text, passwordHash, img)
	if err != nil {
		return "", err
	}
	defer raw.Close()
	for raw.Next() {
		err := raw.Scan(&id)
		if err != nil {
			return "", err
		}
	}
	return id, nil
}

func (t *Thread) CheckPassword(id, passwordHash string) bool {
	var passwordHashInThread string
	raw, err := t.Db.Query(`SELECT password_hash FROM threads WHERE id = $1`, id)
	if err != nil {
		return false
	}
	defer raw.Close()
	for raw.Next() {
		err := raw.Scan(&passwordHashInThread)
		if err != nil {
			return false
		}
	}
	return passwordHashInThread == passwordHash
}

func (t *Thread) GetThreads(limit int32) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, limit)
	var id, head, text, img string
	raw, _ := t.Db.Query(`SELECT id, head, text, img FROM threads WHERE thread_id IS NULL LIMIT $1`, limit)
	defer raw.Close()
	for raw.Next() {
		raw.Scan(&id, &head, &text, &img)
		m := make(map[string]interface{}, 4)
		m["id"] = id
		m["head"] = head
		m["text"] = text
		m["img"] = img
		result = append(result, m)
	}
	return result
}

func (t *Thread) GetThreadAndPosts(threadId string, limit int32) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, limit)
	var id, head, text, img string
	raw, _ := t.Db.Query(`SELECT id, head, text, img FROM threads WHERE id = $1 OR thread_id = $1 ORDER BY id LIMIT $2`, threadId, limit)
	defer raw.Close()
	for raw.Next() {
		raw.Scan(&id, &head, &text, &img)
		m := make(map[string]interface{}, 5)
		m["id"] = id
		m["head"] = head
		m["text"] = text
		m["img"] = img
		result = append(result, m)
	}
	return result
}
