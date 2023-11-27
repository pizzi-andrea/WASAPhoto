package database

func (db *appdbimpl) PostUser(user User) (newU User, _error error) {

	_error = db.c.QueryRow("INSERT INTO Users (username) VALUES (?) RETURNING *", user.Username).Scan(&newU.uid, &newU.Username)
	return
}
