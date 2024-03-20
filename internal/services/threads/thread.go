package service_threads

import (
	"exach/internal/database/models"

	"github.com/jmoiron/sqlx"
)

type ThreadService struct {
	db *sqlx.DB
}

func NewThreadService(db *sqlx.DB) *ThreadService {
	return &ThreadService{
		db: db,
	}
}

func (t *ThreadService) CreateThread(thread models.Thread) (int, error) {
	lastInsertId := 0
	err := t.db.QueryRow(`
	INSERT INTO threads (head, text, password_hash, img) VALUES ($1, $2, $3, $4) RETURNING id
	`, thread.Head, thread.Text, thread.Password_hash, thread.Img).Scan(&lastInsertId)
	return lastInsertId, err
}

func (t *ThreadService) CheckPassword(id, passwordHash string) bool {
	result := models.Thread{}
	t.db.Get(&result, `SELECT password_hash FROM threads WHERE id = $1`, id)
	return result.Password_hash == passwordHash
}

func (t *ThreadService) GetThreads(limit int32) []models.Thread {
	result := []models.Thread{}
	t.db.Select(
		&result,
		`SELECT id, head, text, img FROM threads LIMIT $1`,
		limit,
	)
	return result
}

func (t *ThreadService) GetThreadById(id int) models.Thread {
	result := models.Thread{}

	t.db.Get(&result, `SELECT * FROM threads WHERE id = $1 LIMIT 1`, id)

	return result
}
