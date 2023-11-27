package database

import (
	"bufio"
	"encoding/json"
)

// give user id (uid) and ger correspoding user. If uid not exist the query will be empty
func (db *appdbimpl) GetUser(uid Id) (User, error) {
	var queryUser string
	var user User
	var w *bufio.Writer

	_error := db.c.QueryRow("SELECT * FROM Users WHERE uid = ?", uid).Scan(&queryUser)
	if _error != nil {
		return user, _error
	}
	_error = json.NewEncoder(w).Encode(user)
	return user, _error
}
