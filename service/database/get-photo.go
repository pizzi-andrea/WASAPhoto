package database

/*
 */
func (db *appdbimpl) GetPhoto(id Id) (img *Photo, err error) {
	var photo Photo

	if err = db.c.QueryRow("SELECT * FROM Photos WHERE photoId = ?", id).Scan(&photo.PhotoId, nil, &photo.DescriptionImg, &photo.TimeUpdate); err != nil {
		return

	}

	img = &photo
	return

}
