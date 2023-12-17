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

	checkPostOwnerStatement := "SELECT userid FROM post WHERE id = ?"
	owner := tx.QueryRow(checkPostOwnerStatement, postid)
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

	deletePostStatement := "DELETE FROM post WHERE id = ?"
	_, err = tx.Exec(deletePostStatement, postid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	deletePostLikesStatement := "DELETE FROM like WHERE postid = ?"
	_, err = tx.Exec(deletePostLikesStatement, postid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	deletePostCommentsStatement := "DELETE FROM comment WHERE postid = ?"
	_, err = tx.Exec(deletePostCommentsStatement, postid)
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
