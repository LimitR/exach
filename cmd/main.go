package main

import (
	servers "gochan/internal/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := servers.NewServer()
	s.Run()

}
