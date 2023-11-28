package database

import "fmt"

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateUser(username string) (User, error) {
	var u User
	query_str := fmt.Sprintf("INSERT INTO user (username) VALUES (%s)", username)
	rows, err := db.c.Query(query_str)
	if err != nil {
		return u,err
	}
	for rows.Next() {
		err := rows.Scan(&u.UserID, &u.Username, &u.Feeling, &u.Bio, &u.PictureURL)
		if err != nil {
			return u, err
		}
	}

	return u, nil
}