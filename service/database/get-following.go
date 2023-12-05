package database

import "database/sql"

/*
GetFollowing return list users following by user.
*/
func (db *appdbimpl) GetFollowing(uid Id, username Username, largeSearch bool) (following []User, err error) {
	var rows *sql.Rows
	if largeSearch {
		rows, err = db.c.Query("SELECT uid, username FROM Followers, Users WHERE to_ = uid AND from_ = ? AND username LIKE '%"+username+"%'", uid)
	} else {
		rows, err = db.c.Query(`SELECT uid, username FROM Followers, Users WHERE to_ = uid AND from_ = ? AND username = ?`, uid, username)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if _err := rows.Scan(&uid, &username); _err != nil {
			err = _err
			return
		}

		following = append(following, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
