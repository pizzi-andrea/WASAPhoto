package database

import "database/sql"

func (db *appdbimpl) GetUserFromId(uid Id) (usr *User, err error) {
	var u User
	err = db.c.QueryRow(`
	SELECT * FROM Users 
	WHERE uid = ?`, uid).Scan(&u.Uid, &u.Username)

	if err == nil {
		usr = &u
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return

}
