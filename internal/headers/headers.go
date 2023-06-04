package headers

import (
	"gochan/internal/database/repository"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	Headers *gin.Engine
	Repo    *repository.Repo
}

func NewHeaders(r *gin.Engine, repo *repository.Repo) *Headers {
	return &Headers{
		Headers: r,
		Repo:    repo,
	}
}
