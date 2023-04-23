package posts_controllers

import (
	"gochan/internal/database/posts"
)

func CreatePost(h *posts.Post, text, img, threadId string) error {
	return h.CreatePost(text, img, threadId)
}
