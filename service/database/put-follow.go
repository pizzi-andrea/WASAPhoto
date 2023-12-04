package database

import "database/sql"

// PutLollow allow "from" user to follow other user called "to".
func (db *appdbimpl) PutFollow(from Id, to Id) (r bool, err error) {
	r = false
	err = db.c.QueryRow("INSERT OR IGNORE INTO Followers (from_, to_) VALUES (?, ?) RETURNING *", from, to).Scan(&from, &to)

	if err == nil {
		r = true
		return
	}

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	return
}
