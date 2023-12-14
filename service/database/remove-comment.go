package database

import "errors"

// GetName is an example that shows you how to query data
func (db *appdbimpl) RemoveComment(commentid int, userid int) error {

	// check if comment exists (select 1)
	comment_exists_statement := "SELECT 1 FROM comment WHERE id = ? AND userid = ?"
	var exists int
	err := db.c.QueryRow(comment_exists_statement, commentid, userid).Scan(&exists)
	if err != nil {
		return err
	}
	if exists != 1 {
		return errors.New("comment does not exist")
	}

	// remove comment
	comment_remove_statement := "DELETE FROM comment WHERE id = ? AND userid = ?"

	// follow user
	_, err = db.c.Exec(comment_remove_statement, commentid, userid)
	if err != nil {
		return err
	}

	return nil

}
