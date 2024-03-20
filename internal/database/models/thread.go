package models

type Thread struct {
	Id            int64  `db:"id" json:"id"`
	Text          string `db:"text" json:"text"`
	Head          string `db:"head" json:"head"`
	Password_hash string `db:"password_hash" json:"passwordHash"`
	Img           string `db:"img" json:"img"`
}
