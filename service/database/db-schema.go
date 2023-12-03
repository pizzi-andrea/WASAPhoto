package database

import "go/types"

type Schema struct {
	name  string
	field map[string]types.Const
}

var Users struct {
	name     string
	Uid      string
	Username string
}

var Photos struct {
	PhotoId        string
	Owner          string
	DescriptionImg string
	TimeUpdate     string
}

var Comments struct {
	CommentId  string
	Author     string
	Photo      string
	Text       string
	TimeUpdate string
}
