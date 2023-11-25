package database

import (
	"bufio"
	"encoding/json"
	"fmt"
)

// give user id (uid) and ger correspoding user. If uid not exist the query will be empty
func (db *appdbimpl) GetUser(uid Id) (User, error) {
	var queryUser string
	var user User
	var w *bufio.Writer

	_error := db.c.QueryRow(fmt.Sprintf("SELECT * FROM Users WHERE uid = %d", uid)).Scan(&queryUser)
	json.NewEncoder(w).Encode(user)

	return user, _error
}
