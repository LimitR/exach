package threads_controllers

import (
	"gochan/internal/database/repository"
)

func GetThreadAndPosts(h *repository.Repo, threadId string, limit int32) []repository.Thread {
	return h.Thread().GetThreadAndPosts(threadId, limit)
}
