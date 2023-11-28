package database

import (
	"database/sql"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUsername(username string) (User,bool, error) {
	var u User
	
	var nullableFeeling sql.NullInt64
	var nullableBio sql.NullString
	var nullablePictureURL sql.NullString

	query_str := "SELECT * FROM user WHERE username = '" + username + "';"
	err := db.c.QueryRow(query_str).Scan(&u.UserID, &u.Username, &nullableFeeling, &nullableBio, &nullablePictureURL)
	
	if err == sql.ErrNoRows {
		return u,false, nil // not found but no error
	}
	if err != nil {
		return u,false, err // error
	}
	
	if nullableFeeling.Valid {
		u.Feeling = int(nullableFeeling.Int64)
	}
	if nullableBio.Valid {
		u.Bio = nullableBio.String
	}
	if nullablePictureURL.Valid {
		u.PictureURL = nullablePictureURL.String
	}
	
	return u,true, nil // found and no error
}
