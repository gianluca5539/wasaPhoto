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

type User struct {
	UserID int `json:"userid"`
	Username string `json:"username"`
	Feeling int `json:"feeling"`
	Bio string `json:"bio"`
	Picture string `json:"picture"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserByUsername(username string) (User, bool, error)
	CreateUser(username string) (User, error)
	DeleteAllData() (error)

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

	var err error

	var tableNameUser string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&tableNameUser)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE user (id INTEGER NOT NULL PRIMARY KEY, username TEXT NOT NULL, feeling INT, bio TEXT, picture TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var tableNamePost string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='post';`).Scan(&tableNamePost)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE post (id INTEGER NOT NULL PRIMARY KEY, userid TEXT NOT NULL, picture TEXT NOT NULL, caption TEXT, createdAt DATETIME NOT NULL, likecount INT NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
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
