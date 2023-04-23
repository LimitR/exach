package headers

import (
	"gochan/internal/database/posts"
	"gochan/internal/database/threads"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	Headers *gin.Engine
	post    *posts.Post
	thread  *threads.Thread
}

func NewHeaders(r *gin.Engine, p *posts.Post, t *threads.Thread) *Headers {
	return &Headers{
		Headers: r,
		post:    p,
		thread:  t,
	}
}
