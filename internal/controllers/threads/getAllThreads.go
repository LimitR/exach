package threads_controllers

import (
	"gochan/internal/database/repository"
)

func GetAllThreads(h *repository.Repo, limit int32) []repository.Thread {
	return h.Thread().GetThreads(limit)
}
