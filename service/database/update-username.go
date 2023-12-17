package database

func (db *appdbimpl) UpdateUsername(id int, username string) error {

	// create a new sql statement
	stmt, err := db.c.Prepare("UPDATE user SET username = ? WHERE id = ?")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(username, id)
	if err != nil {
		return err
	}

	return nil
}
