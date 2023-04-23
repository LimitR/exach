package threads_controllers

import (
	"gochan/internal/database/threads"
	"gochan/pkg/auth"
)

func Login(h *threads.Thread, id, password string) bool {
	passwordHash := auth.CreateSum(password)
	return h.CheckPassword(id, passwordHash)
}
