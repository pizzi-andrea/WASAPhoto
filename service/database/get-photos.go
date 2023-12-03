package database

import (
	"database/sql"
	"fmt"
)

/*
Give uid in input and get all photos updated
*/
func (db *appdbimpl) GetPhotos(uid Id) (photos StreamPhotos, err error) {
	var photo Photo
	var id Id
	var rows *sql.Rows
	if rows, err = db.c.Query("SELECT * FROM Photos WHERE owner = ?", uid); err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return

	}

	for rows.Next() {
		if err = rows.Scan(&photo.PhotoId, &id, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); err != nil {
			fmt.Println(fmt.Errorf("%w", err))
			return
		}

		photos = append(photos, photo)

	}
	return

}
