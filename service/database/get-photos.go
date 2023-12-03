package database

import (
	"database/sql"
	"fmt"
)

type OrderBy int

const (
	timeUpdate OrderBy = iota
)

type Ordering int

const (
	asc Ordering = iota
	desc
)

/*
Give uid in input and get all photos updated
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

	if rows, err = db.c.Query("SELECT * FROM Photos WHERE owner = ? ORDER BY ? ?", uid, ordy, oord); err != nil {
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
