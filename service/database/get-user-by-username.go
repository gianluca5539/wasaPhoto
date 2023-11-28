package database

import "database/sql"

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUsername(username string) (User,bool, error) {
	var u User
	query_str := "SELECT * FROM user WHERE username = '" + username + "';"
	err := db.c.QueryRow(query_str).Scan(&u.UserID, &u.Username, &u.Feeling, &u.Bio, &u.PictureURL)
	if err == sql.ErrNoRows {
		return u,false, nil // not found but no error
	}
	if err != nil {
		return u,false, err // error
	}
	return u,true, nil // found and no error
}
