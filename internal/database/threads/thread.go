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

func (t *Thread) CreateThread(head, text, passwordHash, img string) string {
	var id string
	raw, _ := t.Db.Query(`
	INSERT INTO threads (head, text, password_hash, img) VALUES ($1, $2, $3, $4) RETURN id
	`, head, text, passwordHash, img)
	defer raw.Close()
	for raw.Next() {
		err := raw.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
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
