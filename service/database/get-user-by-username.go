package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUsername(string username) (string, error) {
	var u User
	query_str := "SELECT * FROM user AND username = '" + username + "';"
	err := db.c.QueryRow(query_str).Scan(&u.id, &u.username, &u.feeling, &u.bio, &u.picture_url)
	return u, err
}
