package database

import "fmt"

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateUser(string username) (string, error) {
	var u User
	query_str := fmt.Sprintf("INSERT INTO user (username) VALUES (%s)", username)
	err := db.c.Exec(query_str).Scan(&u.id, &u.username, &u.feeling, &u.bio, &u.picture_url)
	return u, err
}