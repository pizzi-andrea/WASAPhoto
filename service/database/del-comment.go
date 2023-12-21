package database

import (
	"database/sql"
	"errors"
)

// DelBan allow to remove ban on user to. When banned is removed user banned will research and watch all information about from user.
// if ban deleted with success function return true and nil error for error if ban not exist function return false and nil value
// if occured error function return false and not-nil error value.
func (db *appdbimpl) DelComment(commentId Id) (r bool, err error) {
	r = false
	var c Comment

	err = db.c.QueryRow("DELETE FROM Comments WHERE commentId = ?", commentId).Scan(&c.CommentId, &c.Author.Uid, &c.Author.Username, commentId, &c.Text, &c.TimeStamp)

	if err == nil {
		r = true
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		r = false
	}
	return
}
