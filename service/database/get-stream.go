package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gianluca5539/WASA/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetStream(ids []int) ([]types.Post, error) {
	// get the "stream" by querying the database for the post table in which "userid" is in the ids array.
	// inner join the user table to get the user's name and profile picture
	// order by the post's timestamp

	var placeholders []string
	var args []interface{}
	for _, id := range ids {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}

	// Join the placeholders with commas and insert them into the query string
	placeholderStr := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT post.id, post.userid, user.username, user.feeling, user.picture, post.picture, post.caption, post.createdAt, post.likeCount FROM post INNER JOIN user ON post.userid = user.id WHERE post.userid IN (%s) ORDER BY post.createdAt DESC", placeholderStr)

	rows, err := db.c.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stream []types.Post
	for rows.Next() {
		var post types.Post
		var nullableLikeCount sql.NullInt64
		var nullableFeeling sql.NullInt64
		var nullableUserPicture sql.NullInt64
		var nullableCaption sql.NullString

		// scan the row into the post
		err = rows.Scan(&post.PostID, &post.UserID, &post.Username, &nullableFeeling, &nullableUserPicture, &post.Picture, &nullableCaption, &post.CreatedAt, &nullableLikeCount)
		if err != nil {
			return nil, err
		}

		// check if the nullable values are null
		if nullableLikeCount.Valid {
			post.LikeCount = int(nullableLikeCount.Int64)
		}
		if nullableFeeling.Valid {
			post.Feeling = int(nullableFeeling.Int64)
		}
		if nullableUserPicture.Valid {
			post.UserPicture = int(nullableUserPicture.Int64)
		}
		if nullableCaption.Valid {
			post.Caption = nullableCaption.String
		}

		// add the post to the list
		stream = append(stream, post)
	}

	// return the list of posts
	return stream, nil
}
