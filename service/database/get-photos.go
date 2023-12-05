package database

import (
	"database/sql"
	"fmt"
)

type OrderBy int

const ( // sort by value
	timeUpdate OrderBy = iota
)

type Ordering int

const ( // Ordering parameters
	asc Ordering = iota
	desc
)

/*
GetPhotos accepts the user ID (uid) as input and extracts all photos uploaded by users. Users can charge up or more
photos (deleted photos will not be listed), in fact the photo stream may be returned empty. If a value of zero is returned
for photos, an error occurred
*/
func (db *appdbimpl) GetPhotos(uid Id, by []OrderBy, ord ...Ordering) (photos StreamPhotos, err error) {
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

	if rows, err = db.c.Query("SELECT * FROM Photos WHERE owner = ? ORDER BY ?, ?", uid, ordy, oord); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&photo.PhotoId, &uid, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); err != nil {
			fmt.Println(fmt.Errorf("%w", err))
			return
		}

		photos = append(photos, photo)

	}
	return

}
