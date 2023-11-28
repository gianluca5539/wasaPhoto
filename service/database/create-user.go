package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) CreateUser(username string) (User, error) {

	var u User

	query_str := "INSERT INTO user (username) VALUES (?);"
	result, err := db.c.Exec(query_str, username)
	if err != nil {
		return u,err
	}

	// Get the last inserted ID
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		// Handle the error
		return u,err
	}

	// create user
	u = User{
		UserID:     int(lastInsertedID),
		Username:   username,
		Feeling:    0,
	}

	
	return u, nil // found and no error
}