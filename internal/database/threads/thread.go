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

func (t *Thread) CreateThread(head, text, passwordHash, img string) {
	t.Db.Exec(`
	INSERT INTO threads (head, text, password_hash, img) VALUES ($1, $2, $3, $4)
	`, head, text, passwordHash, img)
}
