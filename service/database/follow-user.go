package database


// GetName is an example that shows you how to query data
func (db *appdbimpl) FollowUser(id int, followedBy int) ( error) {
	// create a new sql statement
	stmt, err := db.c.Prepare("insert into follow (follow, followedBy) values (?, ?)")
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