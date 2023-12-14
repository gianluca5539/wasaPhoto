package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateNewPost(userID int, image int, caption string, time int) (id int64, error error) {
	// create a new sql statement
	stmt, err := db.c.Prepare("insert into post (userid,picture,caption, createdAt) values (?,?,?,?)")
	if err != nil {
		return -1, err
	}

	// execute the sql statement and get the id of the post
	res, err := stmt.Exec(userID, image, caption, time)
	if err != nil {
		return -1, err
	}

	// get the id of the post
	id, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil

}
