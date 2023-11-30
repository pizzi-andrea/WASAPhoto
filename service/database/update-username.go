package database

func (db *appdbimpl) SetUsername(uid Id, username string) (user User, err error) {

	err = db.c.QueryRow("UPDATE Users SET username=? WHERE uid = ? RETURNING *", username, uid).Scan(&user.Uid, &user.Username)
	return
}
