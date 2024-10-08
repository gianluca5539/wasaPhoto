package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetFollowers(id int) ([]int, error) {
	var ids []int

	// create a new sql statement
	stmt, err := db.c.Prepare("SELECT followedBy FROM follow WHERE follow = ?")
	if err != nil {
		return ids, err
	}

	// execute the sql statement
	rows, err := stmt.Query(id)
	if err != nil {
		return ids, err
	}

	// iterate over each row
	for rows.Next() {
		var id int

		// scan the row into the id
		err = rows.Scan(&id)
		if err != nil {
			return ids, err
		}

		// add the id to the list
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// return the list of ids
	return ids, nil

}
