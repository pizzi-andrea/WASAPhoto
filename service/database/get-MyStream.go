package database

import (
	"database/sql"
	"fmt"
)

func (db *appdbimpl) GetMyStream(uid Id) (photos StreamPhotos, err error) {
	var photo Photo
	var owner Id
	var rows *sql.Rows

	_, err = db.c.Exec(`	
PRAGMA temp_store = 3;
	
CREATE TEMP TABLE UsersNoBanned AS SELECT uid FROM Users u , Bans b 
	WHERE b.to_ != ? AND u.uid = b.from_ ;

CREATE TEMP TABLE MyFollowers AS SELECT uid FROM UsersNoBanned unb, Followers f
	WHERE f.from_ = ? AND unb.uid = f.to_;



CREATE TEMP TABLE MyStream AS SELECT photoId, owner, descriptionImg, imageData, timeUpdate FROM MyFollowers myf, Photos p
	WHERE myf.uid = p.owner;
	`, uid, uid)

	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return
	}

	if rows, err = db.c.Query("SELECT * FROM MyStream"); err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return
	}

	for rows.Next() {
		if _err := rows.Scan(&photo.PhotoId, &owner, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); _err != nil {
			err = _err
			return
		}

		photos = append(photos, photo)

	}
	return

}
