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

	// Start Database
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
	GetUser(uid Id) (User, error)
	GetUsers(username Username) ([]User, error)
	PostUser(user User) (newU User, _error error)
	GetFollower(uid Id) (followers []User, err error)
	GetFollowing(uid Id) (following []User, err error)
	GetMyStream(uid Id) (photos StreamPhotos, err error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

/*
var (
	t_user = Table{
		NameTable: "user",

		Fields: Field{

			"uid": {
				Ftype:      SQLITE_INTEGER,
				PrimaryKey: true,
				Unique:     true,
				NotNull:    true,
			},

			"username": {
				Ftype:      SQLITE_TEXT,
				PrimaryKey: false,
				Unique:     true,
				NotNull:    true,
				Check:      "",
			},
		},
	}

	t_photo = Table{
		NameTable: "photo",

		Fields: Field{

			"photoId": {
				Ftype:      SQLITE_INTEGER,
				PrimaryKey: true,
				Unique:     true,
				NotNull:    true,
			},

			"imageData": {
				Ftype:      SQLITE_BLOB,
				PrimaryKey: false,
				Unique:     false,
				NotNull:    true,
			},

			"timeUpdate": {
				Ftype:      SQLITE_TIME,
				PrimaryKey: false,
				Unique:     false,
				NotNull:    true,
			},

			"descriptionImg": {
				Ftype:      SQLITE_TEXT,
				PrimaryKey: false,
				Unique:     false,
				NotNull:    true,
			},

			"Owner": {
				Ftype:      SQLITE_INTEGER,
				PrimaryKey: false,
				Unique:     false,
				NotNull:    true,
			},
		},
	}
)
*/
// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
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
