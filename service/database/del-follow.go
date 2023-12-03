package database

import "database/sql"

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
