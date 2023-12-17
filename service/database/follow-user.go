package database

import (
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) FollowUser(id int, followedBy int) error {

	follow_statement := "insert into follow (follow, followedBy) values (?, ?)"
	check_follow_statement := "SELECT EXISTS(SELECT 1 FROM follow WHERE follow = ? AND followedBy = ?)"

	// Check if user is already followed and if not follow the user. Do everything in a transaction
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// check if user is already followed
	var exists bool
	err = tx.QueryRow(check_follow_statement, id, followedBy).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user is already followed")
	}

	// follow user
	_, err = tx.Exec(follow_statement, id, followedBy)
	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil

}
