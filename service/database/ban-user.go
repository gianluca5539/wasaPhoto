package database

import (
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(id int, bannedBy int) error {
	ban_statement := "insert into bannedUsers (banned, bannedBy) values (?, ?)"
	remove_follow_statement := "DELETE FROM follow WHERE follow = ? AND followedBy = ?"
	check_ban_statement := "SELECT EXISTS(SELECT 1 FROM bannedUsers WHERE banned = ? AND bannedBy = ?)"

	// Check if user is already banned and if not ban the user. Do everything in a transaction
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if user is already banned
	var exists bool
	err = tx.QueryRow(check_ban_statement, id, bannedBy).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user is already banned")
	}

	// Remove follow
	_, err = tx.Exec(remove_follow_statement, bannedBy, id)
	if err != nil {
		return err
	}

	// Ban user
	_, err = tx.Exec(ban_statement, id, bannedBy)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
