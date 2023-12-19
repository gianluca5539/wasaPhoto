package database

import (
	"database/sql"

	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetComments(postid int) ([]types.Comment, error) {
	query := "SELECT comment.id, comment.userid, comment.postid, user.username, user.picture, user.feeling, comment.text, comment.createdat FROM comment INNER JOIN user ON comment.userid = user.id WHERE postid = ? GROUP BY comment.id ORDER BY createdat ASC"

	rows, err := db.c.Query(query, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []types.Comment
	for rows.Next() {
		var comment types.Comment
		var nullablePicture sql.NullInt64
		var nullableFeeling sql.NullInt64

		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Username, &nullablePicture, &nullableFeeling, &comment.Text, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}

		if nullablePicture.Valid {
			comment.Picture = int(nullablePicture.Int64)
		} else {
			comment.Picture = -1
		}

		if nullableFeeling.Valid {
			comment.Feeling = int(nullableFeeling.Int64)
		} else {
			comment.Feeling = 0
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// return the list of posts
	return comments, nil
}
