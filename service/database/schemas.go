package database

import (
	"container/list"
	"image"
	"regexp"
	"time"
)

type Id = uint64           // Identificator at 64-bit
type TimeStamp = time.Time // this components describe timestamp value conform to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) specification
type Username = string     // username of a user

/*this object rappresent a photo*/
type Photo struct {
	PhotoId        Id
	TimeUpdate     TimeStamp
	ImageData      image.Image // data
	DescriptionImg string      // image description
	MaxLength      int
}

/*this object rappresent a user*/
type User struct {
	Uid      Id
	Username Username
}

// get uid of user
func (usr *User) GetId() Id {
	return usr.Uid
}

// create new user object
func NewUser(uid Id, username Username) User {
	usr := User{
		Uid:      uid,
		Username: username,
	}

	return usr
}

/*user profile rappresentation*/
type Profile struct {
	User      User
	Stream    StreamPhotos
	Follower  int //number user that follow a specific user
	Following int //numer of users following by specific user
}

/*this object rappresent a photo*/
type StreamPhotos struct { //model of stream of photos

	Items    list.List
	MinItems int
	MaxItems int
}

/*this object rappresent a comment on a photo.*/
type Comment struct {
	User      User      // this object rappresent a user
	Text      string    // comment text encoded in UNICODE format
	TimeStamp TimeStamp // this components describe timestamp value conform to RFC3339 specification
}

/* define rule for string. A rule is compose by limit and regex*/
type Rule struct {
	Min        int
	Max        int
	Pattern, _ regexp.Regexp
}

/*
This object rappresent a token identification. The token will be used by users to authenticate to the system.

	The token is composed from two fields:
	  - tokenId: corresponding to uid of owner
	  - owner:   username of owner
*/
type Token struct {
	TokenId Id
	Owner   Username
}
