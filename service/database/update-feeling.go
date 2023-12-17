package database

func (db *appdbimpl) UpdateFeeling(id int, feeling int) error {

	// create a new sql statement
	stmt, err := db.c.Prepare("UPDATE user SET feeling = ? WHERE id = ?")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(feeling, id)
	if err != nil {
		return err
	}

	return nil
}
