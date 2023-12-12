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
	"github.com/gianluca5539/WASA/service/types"
)



// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserByUsername(username string) (types.User, bool, error)
	GetUserByUserID(id int) (types.User, bool, error)
	CreateUser(username string) (types.User, error)
	IsUserBanned(id int, bannedBy int) (bool, error)
	GetFollowers(id int) ([]int, error)
	GetFollowing(id int) ([]int, error)
	UpdateUsername(id int, username string) (error)
	UpdateBio(id int, bio string) (error)
	UpdateFeeling(id int, feeling int) (error)
	UpdatePicture(id int, picture int) (error)
	FollowUser(id int, followedBy int) (error)
	UnFollowUser(id int, followedBy int) (error)
	BanUser(id int, bannedBy int) (error)
	UnBanUser(id int, bannedBy int) (error)
	CreateNewPost(userID int, image int, caption string , time int) (id int64, error error)
	GetPostsByUserID(id int) ([]types.UserPost, error)

	ExecuteSQLDB(code string) ()

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

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
