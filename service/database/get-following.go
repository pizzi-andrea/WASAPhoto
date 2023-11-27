package database

func (db *appdbimpl) GetFollowing(uid Id) (following []User, err error) {
	var username Username
	rows, err := db.c.Query(`SELECT uid, username FROM Followers, Users
	WHERE to_ = uid AND from_ = ?`, uid)

	for rows.Next() {
		if _err := rows.Scan(&uid, &username); _err != nil {
			err = _err
			return
		}

		following = append(following, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
