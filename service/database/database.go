/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	//  Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserFromId(uid Id) (usr *User, err error) //
	GetUserFromUser(Username Username) (usr *User, err error)
	GetUsers(username Username, largeSearch bool) (users []User, err error)
	PostUser(username Username) (usr *User, err error)
	GetFollowers(uid Id, username Username, largeSearch bool) (followers []User, err error)
	GetFollowing(uid Id, username Username, largeSearch bool) (following []User, err error)
	GetMyStream(uid Id, username Username, largeSearch bool, by []OrderBy, ord ...Ordering) (photos StreamPhotos, err error)
	GetPhotoStream(uid, photoId Id) (img *Photo, err error)
	GetBanned(uid Id) (banned []User, err error)
	SetUsername(uid Id, username string) (usr *User, err error)                       //  update username of user associted to uid
	GetPhotos(uid Id, by []OrderBy, ord ...Ordering) (photos StreamPhotos, err error) //  give user id and put all photos posted by user associated to uid
	IsBanned(from Id, to Id) (r bool, err error)
	DelFollow(from, to Id) (r bool, err error)
	IsFollower(from Id, to Id) (r bool, err error)
	PutFollow(from Id, to Id) (r bool, err error)
	PutBan(from, to Id) (r bool, err error)
	DelBan(from, to Id) (r bool, err error)
	PutPhoto(imgData []byte, desc string, owner Id) (photo *Photo, err error)
	GetPhoto(id Id) (img *Photo, err error)
	DelPhoto(id Id) (r bool, err error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	//  Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='WASAPhoto';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {

		if err = initDb(db); err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
