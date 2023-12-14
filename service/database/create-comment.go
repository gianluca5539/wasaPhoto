package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateComment(postid int, userid int, text string) error {

	comment_statement := "insert into comment (postid, userid, text) values (?, ?, ?)"

	// follow user
	_, err := db.c.Exec(comment_statement, postid, userid, text)
	if err != nil {
		return err
	}

	return nil

}
