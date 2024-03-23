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
	INSERT INTO threads (head, text, user_name, img, thread_id) VALUES ($1, $2, $3, $4, $5) RETURNING id
	`, thread.Head, thread.Text, thread.UserName, thread.Img, thread.ThreadId).Scan(&lastInsertId)
	return lastInsertId, err
}

func (t *ThreadService) CheckPassword(id, userName string) bool {
	result := models.Thread{}
	t.db.Get(&result, `SELECT password_hash FROM threads WHERE id = $1`, id)
	return result.UserName == userName
}

func (t *ThreadService) GetThreadsFromBoard(board string, limit int) []models.Thread {
	result := []models.Thread{}
	t.db.Select(
		&result,
		`SELECT * FROM threads WHERE thread_id = '' and board = $1 LIMIT $2`,
		board,
		limit,
	)
	return result
}

func (t *ThreadService) GetThreads(limit int32) []models.Thread {
	result := []models.Thread{}
	t.db.Select(
		&result,
		`SELECT * FROM threads WHERE thread_id = ''  LIMIT $1`,
		limit,
	)
	return result
}

func (t *ThreadService) GetThreadById(id int, board string) models.Thread {
	result := models.Thread{}

	t.db.Get(&result, `SELECT * FROM threads WHERE id = $1 and thread_id = '' and board = $2 or board is null LIMIT 1`, id, board)

	return result
}

func (t *ThreadService) GetCommentsByThreadId(id int64) []models.Thread {
	result := []models.Thread{}

	t.db.Select(
		&result,
		`SELECT * FROM threads where thread_id = $1`,
		id,
	)
	return result
}
