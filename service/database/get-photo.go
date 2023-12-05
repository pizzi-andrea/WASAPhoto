package database

/*
The GetPhoto method allows you to get a photo. If the photo exists it will return the photo otherwise
it will return a value of zero for the photo.
*/
func (db *appdbimpl) GetPhoto(id Id) (img *Photo, err error) {
	var photo Photo

	if err = db.c.QueryRow("SELECT * FROM Photos WHERE photoId = ?", id).Scan(&photo.PhotoId, nil, &photo.DescriptionImg, &photo.TimeUpdate); err != nil {
		return

	}

	img = &photo
	return

}
