package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateComment(postid int, userid int, text string, createdat int) (commentID int, err error) {

	commentStatement := "insert into comment (postid, userid, text, createdat) values (?, ?, ?, ?)"

	// follow user
	res, err := db.c.Exec(commentStatement, postid, userid, text, createdat)
	if err != nil {
		return 0, err
	}

	// return the comment id (last insert id)
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
