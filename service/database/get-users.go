package database

import (
	"database/sql"
)

// get users, Users getted can filtered by  usrname.
// With offset parameters is possible specify the number of rows to skip from the beginning of the table
func (db *appdbimpl) GetUsers(username Username) (users []User, err error) {
	var rows *sql.Rows
	var uid Id
	var name Username

	if username != "" {
		rows, err = db.c.Query("SELECT * FROM Users WHERE username = ?", username)
	} else {
		rows, err = db.c.Query("SELECT * FROM Users")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&uid, &name); err != nil {
			return users, err
		}
		users = append(users, User{
			Uid:      uid,
			Username: name,
		})
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil

}
