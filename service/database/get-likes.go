package database

import (
	"database/sql"

	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetLikes(postid int) ([]types.User, error) {
	query := "SELECT user.id, user.username, user.picture, user.feeling, user.bio, FROM like INNER JOIN user ON like.userid = user.id WHERE like.postid = ? ORDER BY id DESC"

	rows, err := db.c.Query(query, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var u types.User
		var nullableFeeling sql.NullInt64
		var nullableBio sql.NullString
		var nullablePicture sql.NullInt64

		err = rows.Scan(&u.UserID, &u.Username, &nullablePicture, &nullableFeeling, &nullableBio)
		if err != nil {
			return nil, err
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

		users = append(users, u)
	}

	// return the list of posts
	return users, nil
}
