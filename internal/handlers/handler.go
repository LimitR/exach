package handlers

import (
	"gochan/internal/database/repository"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Handlers *gin.Engine
	Repo     *repository.Repo
}

func NewHandlers(r *gin.Engine, repo *repository.Repo) *Handlers {
	return &Handlers{
		Handlers: r,
		Repo:     repo,
	}
}
