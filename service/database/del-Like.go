package database

/*
DelLike allows users to remove likes on the photo.
The operation must be idepotent therefore attempting to delete a like whether it exists or not will always give a positive result.
The function will return unsuccessfully if an error occurs
*/
func (db *appdbimpl) DelLike(uid Id, photoId Id) (r bool, err error) {
	var l *User = nil
	if l, err = db.GetLike(uid, photoId); err != nil {
		return false, err
	}

	if l == nil {
		return false, nil
	}

	if err = db.c.QueryRow("DELETE FROM Likes WHERE user = ? AND photo = ? RETURNING *", uid, photoId).Scan(&uid, &photoId); err != nil {
		return r, err
	}

	r = true
	return

}
