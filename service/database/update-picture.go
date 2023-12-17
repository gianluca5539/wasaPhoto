package database

func (db *appdbimpl) UpdatePicture(id int, picture int) error {

	// create a new sql statement
	stmt, err := db.c.Prepare("UPDATE user SET picture = ? WHERE id = ?")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(picture, id)
	if err != nil {
		return err
	}

	return nil
}
