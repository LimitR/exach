package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db     *sqlx.DB
	post   *Post
	thread *Thread
}

func NewRepo(driver, source string) *Repo {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		log.Fatalln(err)
	}
	repo := &Repo{db: db}

	repo.post = NewPost(db)
	repo.thread = NewThread(db)
	repo.post.createTableOrNotExists()
	repo.thread.createTableOrNotExists()

	return repo
}

func (r *Repo) Post() *Post {
	return r.post
}

func (r *Repo) Thread() *Thread {
	return r.thread
}
