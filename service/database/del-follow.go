package database

import "database/sql"

// DelFollow  allow to delete follow to user. When photo is deleted all comments and like associated will be deleted.
// if photo deleted with success function return true and nil error for error if photo not exist function return false and nil value
// if occured error function return false and not-nil error value.
func (db *appdbimpl) DelFollow(from, to Id) (r bool, err error) {

	err = db.c.QueryRow("DELETE FROM Followers WHERE from_= ? AND to_ = ? ", from, to).Scan(&from, &to)
	if err == nil {
		r = true
		return
	}
	if err == sql.ErrNoRows {
		err = nil
		r = false
	}
	return

}
