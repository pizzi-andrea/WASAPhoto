package database

import (
	"database/sql"
)

/*
Check if user "from" put like to "to"
*/
func (db *appdbimpl) IsFollower(from Id, to Id) (res bool, err error) {

	err = db.c.QueryRow("SELECT * FROM Followers WHERE from_ = ? AND to_ = ?", from, to).Scan(&from, &to)
	if err == sql.ErrNoRows {
		res = false
		err = nil
	} else {
		res = true
	}
	return

}
