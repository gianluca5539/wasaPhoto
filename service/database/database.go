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
	UpdateUsername(id int, username string) error
	UpdateBio(id int, bio string) error
	UpdateFeeling(id int, feeling int) error
	UpdatePicture(id int, picture int) error
	FollowUser(id int, followedBy int) error
	UnFollowUser(id int, followedBy int) error
	BanUser(id int, bannedBy int) error
	UnBanUser(id int, bannedBy int) error
	CreateNewPost(userID int, image int, caption string, time int) (id int64, error error)
	GetPostsByUserID(id int) ([]types.UserPost, error)
	GetStream(ids []int) ([]types.Post, error)
	CreateComment(postid int, userid int, text string, createdat int) (int, error)
	RemoveComment(commentid int, userid int) error
	GetComments(postid int) ([]types.Comment, error)
	GetLikes(postid int) ([]types.User, error)
	LikePost(postid int, userid int) error
	UnLikePost(postid int, userid int) error
	DeletePostCascading(userid int, postid int) error

	ExecuteSQLDB(code string)

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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var user string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE user (id INTEGER NOT NULL PRIMARY KEY, username TEXT, feeling INTEGER, bio TEXT, picture INTEGER);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var comment string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comment';`).Scan(&comment)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comment (id INTEGER NOT NULL PRIMARY KEY, userid INTEGER NOT NULL, postid INTEGER NOT NULL, text TEXT NOT NULL, createdat INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var like string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='like';`).Scan(&like)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE like (id INTEGER NOT NULL PRIMARY KEY, userid INTEGER NOT NULL, postid INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var post string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='post';`).Scan(&post)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE post (id INTEGER NOT NULL PRIMARY KEY, userid INTEGER NOT NULL, picture INT NOT NULL, caption TEXT, createdat INT NOT NULL, likecount INT NOT NULL DEFAULT 0);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var follow string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follow';`).Scan(&follow)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE follow (id INTEGER NOT NULL PRIMARY KEY, follow INTEGER NOT NULL, followedBy INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var ban string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='ban';`).Scan(&ban)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE ban (id INTEGER NOT NULL PRIMARY KEY, banned INTEGER NOT NULL, bannedBy INT NOT NULL);`
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
