package database

func (db *appdbimpl) PutPhoto(imgData []byte, desc string, owner Id) (photo *Photo, err error) {
	var p Photo
	if err = db.c.QueryRow("INSERT INTO Photos(owner, descriptionImg, imageData, timeUpdate) VALUES(?, ?, ?, strftime('%Y-%m-%dT%H:%M:%fZ', 'now')) RETURNING *", owner, desc, imgData).Scan(&p.PhotoId, &owner, &p.DescriptionImg, &p.ImageData, &p.TimeUpdate); err != nil {
		return
	}
	photo = &p
	return
}
