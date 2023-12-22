package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetUserByUserID(id int) (types.User, bool, error) {
	var u types.User

	var nullableFeeling sql.NullInt64
	var nullableBio sql.NullString
	var nullablePicture sql.NullInt64

	query := "SELECT id, username, feeling, bio, picture FROM user WHERE id = ?"
	err := db.c.QueryRow(query, id).Scan(&u.UserID, &u.Username, &nullableFeeling, &nullableBio, &nullablePicture)

	if errors.Is(err,sql.ErrNoRows) {
		fmt.Println("No rows were returned!")
		return u, false, nil // not found but no error
	}
	if err != nil {
		return u, false, err // error
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

	return u, true, nil // found and no error
}
