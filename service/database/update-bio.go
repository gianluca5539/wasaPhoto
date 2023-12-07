package database


func (db *appdbimpl) UpdateBio(id int, bio string) (error) {

	// create a new sql statement
	stmt, err := db.c.Prepare("UPDATE user SET bio = ? WHERE id = ?")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(bio, id)
	if err != nil {
		return err
	}

	return nil
}