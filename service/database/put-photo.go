package database

// PutPhoto is function to permit upload new photo in the system. Actualy All data and metadata will tored in the db file. Large
// file could decrese db performance . If photo is updated with success function will return it otherwise return nil value for photo and not nil value
// for error variable
func (db *appdbimpl) PutPhoto(imgData []byte, desc string, owner Id) (photo *Photo, err error) {
	var p Photo
	if err = db.c.QueryRow("INSERT OR IGNORE INTO Photos(owner, descriptionImg, imageData, timeUpdate) VALUES(?, ?, ?, strftime('%Y-%m-%dT%H:%M:%fZ', 'now')) RETURNING *", owner, desc, imgData).Scan(&p.PhotoId, &owner, &p.DescriptionImg, &p.ImageData, &p.TimeUpdate); err != nil {
		return
	}
	photo = &p
	return
}
