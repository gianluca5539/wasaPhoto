package database

import (
	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateUser(username string) (types.User, error) {

	var u types.User

	query_str := "INSERT INTO user (username) VALUES (?);"
	result, err := db.c.Exec(query_str, username)
	if err != nil {
		return u, err
	}

	// Get the last inserted ID
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		// Handle the error
		return u, err
	}

	// create user
	u = types.User{
		UserID:   int(lastInsertedID),
		Username: username,
		Feeling:  0,
	}

	return u, nil // found and no error
}
