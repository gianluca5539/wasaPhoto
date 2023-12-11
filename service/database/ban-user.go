package database


// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(id int, bannedBy int) ( error) {
	// create a new sql statement
	stmt, err := db.c.Prepare("insert into bannedUsers (banned, bannedBy) values (?, ?)")
	if err != nil {
		return err
	}

	// execute the sql statement
	_, err = stmt.Exec(id, bannedBy)
	if err != nil {
		return err
	}

	return nil

}