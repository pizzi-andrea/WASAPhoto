package database

import (
	"database/sql"
)

/*
GetMyStream return user's stream or part of this. The function allow to search photos in stream by username owner, is largeSearch flag is true
will get all photos of users that have similiar username. If flag is false will return photos of users exactly username gived in input.
*/
func (db *appdbimpl) GetMyStream(uid Id, username Username, largeSearch bool, by []OrderBy, ord ...Ordering) (photos []Post, err error) {
	var photo Photo
	var c []Comment
	var likes []User
	var rows *sql.Rows
	var ordy string
	var oord string
	if len(by) == 0 {
		by = append(by, OrderBy(timeUpdate))
	}
	if len(ord) == 0 {
		ord = append(ord, Ordering(desc))
	}

	switch ord[0] {
	case asc:
		oord = "ASC"
	case desc:
		oord = "DESC"
	default:
		oord = "DESC"
	}

	switch by[0] {
	case timeUpdate:
		ordy = "timeUpdate"

	}
	if largeSearch {
		if _, err = db.c.Exec(`	
			PRAGMA temp_store = 3;
				
			CREATE TEMP TABLE IF NOT EXISTS FollowersUser AS
			SELECT uid 
			FROM Followers
			WHERE from_ = ? AND to_ = uid AND username LIKE'%`+username+`%';
			`, uid); err != nil {
			return nil, err
		}

	} else {
		if _, err = db.c.Exec(`	
			PRAGMA temp_store = 3;
			CREATE TEMP TABLE IF NOT EXISTS FollowersUser AS
			SELECT uid 
			FROM Followers
			WHERE from_ = ? AND to_ = uid AND username LIKE `+username+`;
			`, uid); err != nil {
			return nil, err
		}
	}
	_, err = db.c.Exec(`
	CREATE TEMP TABLE IF NOT EXISTS PhotoStream AS
	SELECT photoId, owner, descriptionImg, imageData, timeUpdate
	FROM FollowersUser, Photos
	WHERE owner = uid
	ORDER BY ? ?;
	`, ordy, oord)

	if err != nil {
		return
	}

	if rows, err = db.c.Query("SELECT * FROM PhotoStream;"); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&photo.PhotoId, &uid, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); err != nil {
			return
		}

		if c, err = db.GetComments(photo.PhotoId, username, true); err != nil {
			return
		}

		if likes, err = db.GetLikes(photo.PhotoId); err != nil {
			return
		}

		photos = append(photos, Post{
			Photo:    photo,
			Likes:    likes,
			Comments: c,
		})

	}

	_, err = db.c.Exec("DROP TABLE FollowersUser; DROP TABLE  MyFollowers; DROP TABLE PhotoStream;")
	return

}
