package database

import (
	"database/sql"
	"github.com/gianluca5539/WASA/service/types"
)


// GetName is an example that shows you how to query data
func (db *appdbimpl) GetPostsByUserID(id int) ([]types.UserPost, error) {
	var posts []types.UserPost

	var nullableLikeCount sql.NullInt64

	// create a new sql statement
	stmt, err := db.c.Prepare("SELECT * FROM post WHERE userid = ? ORDER BY createdAt DESC")
	if err != nil {
		return []types.UserPost{}, err
	}

	// execute the sql statement
	rows, err := stmt.Query(id)
	if err != nil {
		return []types.UserPost{}, err
	}

	// iterate over each row
	for rows.Next() {
		var p types.UserPost
		var userid int

		// scan the row into the variables
		err := rows.Scan(&p.PostID, &userid, &p.Picture, &p.Caption,  &p.CreatedAt, &nullableLikeCount)
		if err != nil {
			return []types.UserPost{}, err
		}
		if nullableLikeCount.Valid {
			p.LikeCount = int(nullableLikeCount.Int64)
		} else {
			p.LikeCount = 0
		}

		// add the post to the list
		posts = append(posts, p)
	}


	return posts, nil

}