package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db     *sqlx.DB
	ctx    context.Context
	cancel context.CancelFunc
	post   *Post
	thread *Thread
}

func NewRepo(driver, source string) *Repo {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		log.Fatalln(err)
	}
	repo := &Repo{db: db}
	ctx, cancel := context.WithCancel(context.Background())

	repo.post = NewPost(db, ctx)
	repo.thread = NewThread(db, ctx)
	repo.ctx = ctx
	repo.cancel = cancel

	repo.post.createTableOrNotExists()
	repo.thread.createTableOrNotExists()
	return repo
}

func (r *Repo) CancelContext() {
	r.cancel()
}

func (r *Repo) Post() *Post {
	return r.post
}

func (r *Repo) Thread() *Thread {
	return r.thread
}
