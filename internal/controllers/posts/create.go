package posts_controllers

import (
	"gochan/internal/database/repository"
)

func CreatePost(h *repository.Repo, text, img, threadId string) error {
	return h.Post().CreatePost(text, img, threadId)
}
