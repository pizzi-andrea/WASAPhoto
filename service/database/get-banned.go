package database

import (
	"database/sql"
)

// GetBanned get the list of all users banned from "uid" user. If function term. with success return no.nil array of banned users
// if occurent error while query working will return nil followers array and error value, id error occured while perform insert
// values in the list will returning the list whit parzial values and error value.
func (db *appdbimpl) GetBanned(uid Id) (followers []User, err error) {
	var username Username
	var rows *sql.Rows

	if rows, err = db.c.Query(`SELECT uid, username FROM Users u, Bans b 
	WHERE b.from_ = ? AND u.uid = b.to_`, uid); err != nil {
		return
	}

	defer rows.Close()

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
