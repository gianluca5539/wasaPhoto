package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) UnFollowUser(id int, followedBy int) ( error) {
	// create a new sql statement
	stmt, err := db.c.Prepare("delete from follow where follow = ? and followedBy = ?")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(id, followedBy)
	if err != nil {
		return err
	}

	return nil

}