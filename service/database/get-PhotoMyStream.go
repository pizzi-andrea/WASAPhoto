package database

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
