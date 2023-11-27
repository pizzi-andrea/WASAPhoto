package database

func (db *appdbimpl) PostUser(user User) (_error error) {

	_, _error = db.c.Exec("INSERT INTO Users (username) VALUES (?)", user.Username)
	return
}
