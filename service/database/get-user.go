package database

// give user id (uid) and ger correspoding user. If uid not exist the query will be empty
func (db *appdbimpl) GetUser(uid Id) (User, error) {
	var queryUser User
	var user User

	_error := db.c.QueryRow("SELECT * FROM Users WHERE uid = ?", uid).Scan(&queryUser.Uid, &queryUser.Username)
	return user, _error
}
