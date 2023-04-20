package headers

import (
	"gochan/internal/database/posts"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	Headers *gin.Engine
	post    *posts.Post
}

func NewHeaders(r *gin.Engine, p *posts.Post) *Headers {
	return &Headers{
		Headers: r,
		post:    p,
	}
}
