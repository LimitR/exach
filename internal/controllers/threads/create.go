package threads_controllers

import (
	"gochan/internal/database/repository"
	"gochan/pkg/auth"
)

func CreateThread(h *repository.Repo, head, text, password, img string) (int64, error) {
	passwordHash := auth.CreateSum(password)
	id, err := h.Thread().CreateThread(head, text, passwordHash, img)
	return id, err
}
