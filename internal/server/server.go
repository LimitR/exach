package servers

import (
	"fmt"
	"log"

	"gochan/internal/database/repository"
	"gochan/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		secret["DB_HOST"], 5432, secret["POSTGRES_USER"], secret["POSTGRES_PASSWORD"], secret["POSTGRES_DB"])
	repo := repository.NewRepo(secret["DB_DRIVER"], dsn)
	h := handlers.NewHandlers(r, repo)
	h.AddHeadersPost()
	h.AddHeadersThreads()
	h.Handlers.Run(secret["HOST"] + ":" + secret["PORT"])
}
