package database

import (
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeletePostCascading(userid int, postid int) error {

	// Check if user is already banned and if not ban the user. Do everything in a transaction
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	check_post_owner_statement := "SELECT userid FROM post WHERE id = ?"
	owner := tx.QueryRow(check_post_owner_statement, postid)
	var ownerid int
	err = owner.Scan(&ownerid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	if ownerid != userid {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return errors.New("you are not the owner of the post")
	}

	delete_post_statement := "DELETE FROM post WHERE id = ?"
	_, err = tx.Exec(delete_post_statement, postid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	delete_post_likes_statement := "DELETE FROM like WHERE postid = ?"
	_, err = tx.Exec(delete_post_likes_statement, postid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	delete_post_comments_statement := "DELETE FROM comment WHERE postid = ?"
	_, err = tx.Exec(delete_post_comments_statement, postid)
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
