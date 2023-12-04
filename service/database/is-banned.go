package database

import (
	"database/sql"
)

/*
check if user "from" banned  user "to"
*/
func (db *appdbimpl) IsBanned(from Id, to Id) (res bool, err error) {
	res = false
	err = db.c.QueryRow("SELECT * FROM Bans WHERE from_ = ? AND to_ = ?", from, to).Scan(&from, &to)
	if err == sql.ErrNoRows {
		err = nil
	} else {
		res = true
	}
	return

}
