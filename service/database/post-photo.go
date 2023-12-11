package database

import (
	"bytes"
	"encoding/base64"
	"time"
)

// PutPhoto is function to permit upload new photo in the system. Actualy All data and metadata will tored in the db file. Large
// file could decrese db performance . If photo is updated with success function will return it otherwise return nil value for photo and not nil value
// for error variable
func (db *appdbimpl) PostPhoto(imgData []byte, desc string, owner Id) (photo *Photo, err error) {
	var p Photo
	base64.NewEncoder(base64.RawStdEncoding, bytes.NewBuffer(imgData))

	if err = db.c.QueryRow("INSERT OR IGNORE INTO Photos(owner, descriptionImg, imageData, timeUpdate) VALUES(?, ?, ?, ?) RETURNING *", owner, desc, imgData, time.Now().Format(time.RFC3339)).Scan(&p.PhotoId, &owner, &p.DescriptionImg, &p.ImageData, &p.TimeUpdate); err != nil {
		return
	}
	photo = &p
	return
}
