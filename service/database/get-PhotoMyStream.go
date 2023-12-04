package database

// GetPhotoStream return a single photo from personal user stream. If photo not belong to the stream will return
// null img pointer and nil error value. If photo is in stream will return it. If occurend an error will return nil pointer and
// error value
func (db *appdbimpl) GetPhotoStream(uid, photoId Id) (img *Photo, err error) {
	var photo StreamPhotos

	if photo, err = db.GetMyStream(uid, "", true, []OrderBy{}); err != nil {
		return nil, err
	}

	for _, p := range photo {
		if p.PhotoId == photoId {
			img = &p
			return
		}

	}

	return

}
