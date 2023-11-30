package database

import (
	"database/sql"
)

func (db *appdbimpl) GetBan(uid Id) (followers []User, err error) {
	var username Username
	var rows *sql.Rows

	if rows, err = db.c.Query(`SELECT uid, username FROM Users u, Bans b 
	WHERE b.to_ = ? AND u.uid = b.from_`, uid); err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&uid, &username); err != nil {
			return
		}

		followers = append(followers, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
