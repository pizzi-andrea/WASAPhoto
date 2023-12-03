package database

import "database/sql"

// GetUserFromUser give username of user and get user if exist or nil if not exist
func (db *appdbimpl) GetUserFromUser(username Username) (usr *User, err error) {
	var u User
	err = db.c.QueryRow(`
	SELECT * FROM Users 
	WHERE username = ?`, username).Scan(&u.Uid, &u.Username)

	if err == nil {
		usr = &u
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
