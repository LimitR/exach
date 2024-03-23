package models

type Thread struct {
	Id       int64  `db:"id" json:"id"`
	Text     string `db:"text" json:"text"`
	Head     string `db:"head" json:"head"`
	UserName string `db:"user_name" json:"user_name"`
	Img      string `db:"img" json:"img"`
	ThreadId string `db:"thread_id"`
	Board    string `db:"board"`
}
