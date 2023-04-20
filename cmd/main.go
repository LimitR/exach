package main

import (
	"gochan/internal/database/posts"
	"gochan/internal/headers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	p := posts.NewPost()
	h := headers.NewHeaders(r, p)
	h.AddPost()
	h.Headers.Run("localhost:5000")
}
