package database

import (
	"database/sql"
)

/*
GetMyStream return user's stream or part of this. The function allow to search photos in stream by username owner, is largeSearch flag is true
will get all photos of users that have similiar username. If flag is false will return photos of users exactly username gived in input.
*/
func (db *appdbimpl) GetMyStream(uid Id, username Username, largeSearch bool, by []OrderBy, ord ...Ordering) (photos StreamPhotos, err error) {
	var photo Photo
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
				
			CREATE TEMP TABLE IF NOT EXISTS UsersNoBanned AS 
			SELECT uid 
			FROM Users u , Bans b 
			WHERE b.to_ != ? AND u.uid = b.from_  AND b.from AND + LIKE '%` + username + `%';
			`); err != nil {
			return nil, err
		}

	} else {
		if _, err = db.c.Exec(`	
			PRAGMA temp_store = 3;
				
			CREATE TEMP TABLE IF NOT EXISTS UsersNoBanned AS 
			SELECT uid 
			FROM Users u , Bans b 
			WHERE b.to_ != ? AND u.uid = b.from_  AND b.from AND + LIKE ` + username + `;
			`); err != nil {
			return nil, err
		}
	}
	_, err = db.c.Exec(`	
CREATE TEMP TABLE IF NOT EXISTS MyFollowers AS 
SELECT uid 
FROM UsersNoBanned unb, Followers f
WHERE f.from_ = ? AND unb.uid = f.to_;



CREATE TEMP TABLE IF NOT EXISTS MyStream AS 
SELECT photoId, owner, descriptionImg, imageData, timeUpdate 
FROM MyFollowers myf, Photos p
WHERE myf.uid = p.owner;
	`, uid)

	if err != nil {
		return
	}

	if rows, err = db.c.Query("SELECT * FROM MyStream ORDER BY ? ?", ordy, oord); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&photo.PhotoId, &uid, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); err != nil {
			return
		}

		photos = append(photos, photo)

	}

	_, err = db.c.Exec("DROP TABLE UsersNoBanned; DROP TABLE  MyFollowers; DROP TABLE MyStream ")
	return

}
