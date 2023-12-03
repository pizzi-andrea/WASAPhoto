package database

func (db *appdbimpl) PostUser(username Username) (usr *User, err error) {
	var u User
	err = db.c.QueryRow("INSERT INTO Users (username) VALUES (?) RETURNING *", username).Scan(&u.Uid, &u.Username)
	if err == nil {
		usr = &u
	}
	return
}
