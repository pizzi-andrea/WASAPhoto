package database

func (db *appdbimpl) GetLike(uid Id, photoId Id) (like *User, err error) {
	var u User

	if err = db.c.QueryRow("SELECT uid, username FROM Likes, Users WHERE user = ? AND photo = ?", uid, photoId).Scan(&u.Uid, &u.Username); err != nil {
		return nil, err
	}
	like = &u
	return like, err

}
