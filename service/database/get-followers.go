package database

import "database/sql"

/*
 */
func (db *appdbimpl) GetFollowers(uid Id, username Username, largeSearch bool) (followers []User, err error) {
	var rows *sql.Rows

	if largeSearch {
		rows, err = db.c.Query("SELECT uid, username FROM Followers, Users WHERE to_ = ? AND from_ = uid AND username LIKE '%"+username+"%'", uid)

	} else {
		rows, err = db.c.Query(`SELECT uid, username FROM Followers, Users WHERE to_ = ? AND from_ = uid AND username = ?`, uid, username)

	}

	if err != nil {
		return nil, err
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
