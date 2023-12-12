package database

import (
	"database/sql"
	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUsername(username string) (types.User,bool, error) {
	var u types.User
	
	var nullableFeeling sql.NullInt64
	var nullableBio sql.NullString
	var nullablePicture sql.NullInt64

	query_str := "SELECT * FROM user WHERE username = '" + username + "';"
	err := db.c.QueryRow(query_str).Scan(&u.UserID, &u.Username, &nullableFeeling, &nullableBio, &nullablePicture)
	
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
		u.Picture = int(nullablePicture.Int64)
	}
	
	return u,true, nil // found and no error
}
