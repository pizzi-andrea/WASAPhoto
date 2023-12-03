package database

import "database/sql"

func (db *appdbimpl) DelPhoto(id Id) (r bool, err error) {
	r = false
	err = db.c.QueryRow("DELETE FROM Photos WHERE  photoId = ? ", id).Err()

	if err == nil {
		r = true
		return
	}
	if err == sql.ErrNoRows {
		err = nil
		r = false
	}
	return
}
