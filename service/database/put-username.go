package database

func (db *appdbimpl) SetUsername(uid Id, username string) (usr *User, err error) {
	var u User
	err = db.c.QueryRow("UPDATE Users SET username=? WHERE uid = ? RETURNING *", username, uid).Scan(&u.Uid, &u.Username)
	if err == nil {
		usr = &u
	}

	return

}
