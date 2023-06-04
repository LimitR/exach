package threads_controllers

import (
	"gochan/internal/database/repository"
	"gochan/pkg/auth"
)

func Login(h *repository.Repo, id, password string) bool {
	passwordHash := auth.CreateSum(password)
	return h.Thread().CheckPassword(id, passwordHash)
}
