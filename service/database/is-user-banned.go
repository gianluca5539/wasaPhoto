package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) IsUserBanned(id int, bannedBy int) (bool, error) {

	// create a new sql statement
	stmt, err := db.c.Prepare("SELECT * FROM bannedUsers WHERE banned = ? AND bannedBy = ?")
	if err != nil {
		return false, err
	}

	// execute the sql statement
	rows, err := stmt.Query(id, bannedBy)
	if err != nil {
		return false, err
	}

	// close the sql statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		return true, nil
	}

	// return the result
	return false, nil
}
