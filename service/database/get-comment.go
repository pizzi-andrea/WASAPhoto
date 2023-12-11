package database

import "database/sql"

func (db *appdbimpl) GetComment(commentId Id) (comment *Comment, err error) {
	var c Comment

	err = db.c.QueryRow("SELECT * FROM Comments WHERE commentId").Scan(&c.CommentId, &c.Author, &commentId, &c.Text, &c.TimeStamp)

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	comment = &c
	return

}
