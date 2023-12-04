package database

import (
	"database/sql"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUserID(id int) (User,bool, error) {
	var u User
	
	var nullableFeeling sql.NullInt64
	var nullableBio sql.NullString
	var nullablePicture sql.NullString

	query_str := "SELECT * FROM user WHERE id = ?"
	err := db.c.QueryRow(query_str, id).Scan(&u.UserID, &u.Username, &nullableFeeling, &nullableBio, &nullablePicture)
	
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
	if nullablePicture.Valid {
		u.Picture = nullablePicture.String
	}
	
	return u,true, nil // found and no error
}