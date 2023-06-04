package servers

import (
	"log"

	"gochan/internal/database/repository"
	"gochan/internal/headers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	secret, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	repo := repository.NewRepo(secret["DB_DRIVER"], secret["DATA_SOURCE_NAME"])
	h := headers.NewHeaders(r, repo)
	h.AddHeadersPost()
	h.AddHeadersThreads()
	h.Headers.Run(secret["HOST"] + ":" + secret["PORT"])
}
