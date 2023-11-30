package database

func (db *appdbimpl) PostUser(user User) (newU User, err error) {

	err = db.c.QueryRow("INSERT INTO Users (username) VALUES (?) RETURNING *", user.Username).Scan(&newU.Uid, &newU.Username)
	return
}
