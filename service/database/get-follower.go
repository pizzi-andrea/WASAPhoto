package database

func (db *appdbimpl) GetFollower(uid Id) (followers []User, err error) {
	var username Username
	rows, err := db.c.Query(`SELECT uid, username FROM Followers, Users
	WHERE to_ = ? AND from_ = uid`, uid)

	for rows.Next() {
		if _err := rows.Scan(&uid, &username); _err != nil {
			err = _err
			return
		}

		followers = append(followers, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
