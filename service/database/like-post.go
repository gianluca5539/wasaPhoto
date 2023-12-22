package database

import (
	"database/sql"
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) LikePost(postid int, userid int) error {

	// create a transaction to insert into like and update post.likeCount
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// check if like exists (select 1)
	likeExistsStatement := "SELECT 1 FROM like WHERE postid = ? AND userid = ?"
	var exists int
	err = tx.QueryRow(likeExistsStatement, postid, userid).Scan(&exists)
	if err != nil && !errors.Is(err,sql.ErrNoRows) {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}
	if exists == 1 {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return errors.New("like already exists")
	}

	// insert like
	likeInsertStatement := "INSERT INTO like (postid, userid) VALUES (?, ?)"
	_, err = tx.Exec(likeInsertStatement, postid, userid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}

	// update post.likeCount
	postUpdateStatement := "UPDATE post SET likecount = likecount + 1 WHERE id = ?"
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
