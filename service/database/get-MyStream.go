package database

func (db *appdbimpl) GetMyStream(uid Id) (photos StreamPhotos, err error) {
	var photo Photo
	var owner Id

	rows, err := db.c.Query(`SELECT * FROM Photos
	WHERE owner = ?`, uid)

	for rows.Next() {
		if _err := rows.Scan(&photo.PhotoId, &owner, &photo.DescriptionImg, &photo.ImageData, &photo.TimeUpdate); _err != nil {
			err = _err
			return
		}

		photos.Items.PushFront(photo)

	}
	return

}
