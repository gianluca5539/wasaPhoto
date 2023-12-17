package database

import (
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) FollowUser(id int, followedBy int) error {

	followStatement := "insert into follow (follow, followedBy) values (?, ?)"
	checkFollowStatement := "SELECT EXISTS(SELECT 1 FROM follow WHERE follow = ? AND followedBy = ?)"

	// Check if user is already followed and if not follow the user. Do everything in a transaction
	tx, err := db.c.Begin()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	// check if user is already followed
	var exists bool
	err = tx.QueryRow(checkFollowStatement, id, followedBy).Scan(&exists)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	if exists {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return errors.New("user is already followed")
	}

	// follow user
	_, err = tx.Exec(followStatement, id, followedBy)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	err = tx.Commit()

	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	return nil

}
