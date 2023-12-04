package database

import "database/sql"

// DelPhoto allow to delete photo update in the system. When photo is deleted all comments and like associated will be deleted.
// if photo deleted with success function return true and nil error for error if photo not exist function return false and nil value
// if occured error function return false and not-nil error value.
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
