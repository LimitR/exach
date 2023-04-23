package threads_controllers

import (
	"gochan/internal/database/threads"
	"gochan/pkg/auth"
)

func CreateThread(h *threads.Thread, head, text, password, img string) (string, error) {
	passwordHash := auth.CreateSum(password)
	id, err := h.CreateThread(head, text, passwordHash, img)
	return id, err
}
