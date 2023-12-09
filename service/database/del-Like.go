package database

import (
	"database/sql"
	"errors"
)

/*
DelLike allows users to remove likes on the photo.
The operation must be idepotent therefore attempting to delete a like whether it exists or not will always give a positive result.
The function will return unsuccessfully if an error occurs
*/
func (db *appdbimpl) DelLike(uid Id, photoId Id) (r bool, err error) {

	if err = db.c.QueryRow("DELETE OR IGNORE Likes WHERE user = ? AND photo = ? RETURNING *", uid, photoId).Scan(&uid, &photoId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
			r = true
			return

		} else {
			r = false
		}
		return
	}

	r = true
	return

}
