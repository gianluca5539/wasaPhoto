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
	like_exists_statement := "SELECT 1 FROM like WHERE postid = ? AND userid = ?"
	var exists int
	err = tx.QueryRow(like_exists_statement, postid, userid).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}
	if exists == 0 || err == sql.ErrNoRows {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return errors.New("like does not exist")
	}

	// remove like
	like_remove_statement := "DELETE FROM like WHERE postid = ? AND userid = ?"
	_, err = tx.Exec(like_remove_statement, postid, userid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}

	// update post.likeCount
	post_update_statement := "UPDATE post SET likecount = likecount - 1 WHERE id = ?"
	_, err = tx.Exec(post_update_statement, postid)
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
