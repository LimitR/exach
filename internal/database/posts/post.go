package posts

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

type Post struct {
	Db *sql.DB
}

func NewPost() *Post {
	secret, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open(secret["DB_DRIVER"], secret["DATA_SOURCE_NAME"])
	if err != nil {
		log.Fatalln(err)
	}
	p := &Post{
		Db: db,
	}
	p.createTableOrNotExists()
	return p
}

func (t *Post) createTableOrNotExists() {
	t.Db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		head TEXT,
		text TEXT NOT NULL,
		thread_id TEXT,
		password_hash TEXT,
		img TEXT
		);`)
}

func (t *Post) CreatePost(text, img, thread_id string) error {
	_, err := t.Db.Exec(`
	INSERT INTO threads (text, img, thread_id) VALUES ($1, $2, $3)
	`, text, img, thread_id)
	return err
}
