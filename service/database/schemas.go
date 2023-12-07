package database

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"regexp"
	"time"
)

type Id = uint64           //   Identificator at 64-bit
type TimeStamp = time.Time //   this components describe timestamp value conform to [RFC3339](https://  datatracker.ietf.org/doc/html/rfc3339#section-5.6) specification
type Username = string     //   username of a user

type Validator interface {
	Verify() bool
}

/*this object rappresent a photo*/
type Photo struct {
	PhotoId        Id
	TimeUpdate     TimeStamp
	ImageData      []byte //   data
	DescriptionImg string //   image description
}

/*this object rappresent a user*/
type User struct {
	Uid      Id
	Username Username
}

/*user profile rappresentation*/
type Profile struct {
	User      User
	Stream    StreamPhotos
	Follower  int //  number user that follow a specific user
	Following int //  numer of users following by specific user

}

/*this object rappresent a photo*/
type StreamPhotos = []Photo //  model of stream of photos

/*this object rappresent a comment on a photo.*/
type Comment struct {
	CommentId Id
	Author    User //  this object rappresent a user
	Photo     Id
	Text      string    //  comment text encoded in UNICODE format
	TimeStamp TimeStamp //  this components describe timestamp value conform to RFC3339 specification

}

func ValidateId(v Id) bool {
	return true
}

func ValidateTimeStamp(time string) bool {
	s, err := regexp.MatchString(" '[1-9]\\d{3}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$", time)
	return 20 <= len(time) && len(time) >= 21 && s && err == nil
}

func ValidateUsername(u string) bool {
	s, err := regexp.MatchString("^.*?$", u)
	return len(u) >= 3 && len(u) <= 16 && s && err == nil
}

func ValidateStream(s StreamPhotos) bool {

	r := len(s) <= 100
	if r {
		for _, p := range s {
			if !p.Verify() {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (p *Photo) Verify() bool {
	r, err := regexp.MatchString("^.*?$", p.DescriptionImg)
	return len(p.ImageData) <= 5000000 &&
		p.GetImg() == nil &&
		ValidateTimeStamp(p.TimeUpdate.Local().Format(time.RFC3339)) &&
		ValidateId(p.PhotoId) &&
		len(p.DescriptionImg) >= 1 &&
		len(p.DescriptionImg) <= 250 &&
		r && err == nil
}

func (p *Photo) GetImg() (img *image.Image) {
	i := bytes.NewBuffer(p.ImageData)
	if pn, err := png.Decode(i); err == nil {
		img = &pn
	} else {
		if jp, err := jpeg.Decode(i); err == nil {
			img = &jp
		}
	}

	return
}

func (u *User) Verify() bool {
	return ValidateUsername(u.Username) && ValidateId(u.Uid)
}

func (p *Profile) Verify() bool {
	return p.User.Verify() && p.Follower >= 0 && p.Following >= 0 && ValidateStream(p.Stream)
}

func (c *Comment) Verify() bool {
	r, err := regexp.MatchString("^.*$", c.Text)
	return c.Author.Verify() && ValidateTimeStamp(c.TimeStamp.Format(time.RFC3339)) && len(c.Text) >= 1 && len(c.Text) <= 250 && r && err == nil
}
