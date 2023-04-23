package threads_controllers

import "gochan/internal/database/threads"

func CreateThread(h *threads.Thread, head, text, passwordHash, img string) (string, error) {
	id, err := h.CreateThread(head, text, passwordHash, img)
	return id, err
}
