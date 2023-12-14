package database

import "database/sql"

func (db *appdbimpl) GetComments(photoId Id, username Username, largeSearch bool) (comments []Comment, err error) {
	var rows *sql.Rows
	var c Comment

	if largeSearch {
		rows, err = db.c.Query("SELECT commentId, author, photo, text_, timeStamp_,  username FROM Users u, Comments c WHERE u.uid = c.author AND  c.photo = ? AND  u.username LIKE '%"+username+"%'", photoId)
	} else {
		rows, err = db.c.Query("SELECT commentId, author, photo, text_, timeStamp_, username FROM Users u, Comments c  WHERE u.uid = c.author AND  c.photo = ? AND u.username = ?", photoId, username)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return comments, err
		}

		if err := rows.Scan(&c.CommentId, &c.Author.Uid, &photoId, &c.Text, &c.TimeStamp, &c.Author.Username); err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}

	err = rows.Err()
	return

}
