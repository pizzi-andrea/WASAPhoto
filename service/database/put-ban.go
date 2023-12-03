package database

import "database/sql"

/*
user "from" banned user "to"
*/
func (db *appdbimpl) PutBan(from, to Id) (r bool, err error) {
	r = false
	err = db.c.QueryRow("INSERT INTO BANS (from_, to_) VALUES (?, ?) RETURNING *", from, to).Scan(&from, &to)

	if err == nil {
		r = true
		return
	}

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	return

}
