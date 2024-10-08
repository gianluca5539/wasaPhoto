package database

import (
	"database/sql"
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) UnLikePost(postid int, userid int) error {

	// create a transaction to insert into like and update post.likeCount
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// check if like exists (select 1)
	likeExistsStatement := "SELECT 1 FROM like WHERE postid = ? AND userid = ?"
	var exists int
	err = tx.QueryRow(likeExistsStatement, postid, userid).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}
	if exists == 0 || errors.Is(err, sql.ErrNoRows) {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return errors.New("like does not exist")
	}

	// remove like
	likeRemoveStatement := "DELETE FROM like WHERE postid = ? AND userid = ?"
	_, err = tx.Exec(likeRemoveStatement, postid, userid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}

	// update post.likeCount
	postUpdateStatement := "UPDATE post SET likecount = likecount - 1 WHERE id = ?"
	_, err = tx.Exec(postUpdateStatement, postid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
