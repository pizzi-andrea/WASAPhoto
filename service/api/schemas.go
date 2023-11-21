package api

import (
	"container/list"
	"image"
	"regexp"
	"time"
)

type Id = uint64           // Identificator at 64-bit
type TimeStamp = time.Time // this components describe timestamp value
type Username = string     // username of a user

// conform to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) specification
type Photo struct { //this object rappresent a photo
	photoId        Id
	timeUpdate     TimeStamp
	imageData      image.Image // data
	descriptionImg string      // image description
	maxLength      int
}

/*this object rappresent a user*/
type User struct {
	uid      Id
	username Username
}

/*user profile rappresentation*/
type Profile struct {
	user      User
	stream    StreamPhotos
	follower  int //number user that follow a specific user
	following int //numer of users following by specific user
}

/*this object rappresent a photo*/
type StreamPhotos struct { //model of stream of photos

	items    list.List
	minItems int
	maxItems int
}

/*this object rappresent a comment on a photo.*/
type Comment struct {
	user      User      // this object rappresent a user
	text      string    // comment text encoded in UNICODE format
	timeStamp TimeStamp // this components describe timestamp value conform to RFC3339 specification
}

/* define rule for string. A rule is compose by limit and regex*/
type Rule struct {
	min        int
	max        int
	pattern, _ regexp.Regexp
}

/*
func (rg *Rule) validate(chr string) bool {
	return len(chr) <= rg.max && len(chr) >= rg.min && rg.pattern.MatchString(chr)
}
*/
