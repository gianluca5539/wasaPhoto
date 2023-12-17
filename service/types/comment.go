// Package types provides the data types for the service.
package types

// Comment represents a comment on a post.
type Comment struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid"`
	PostID    int    `json:"postid"`
	Username  string `json:"username"`
	Picture   int    `json:"picture"`
	Feeling   int    `json:"feeling"`
	Text      string `json:"text"`
	CreatedAt int    `json:"createdat"`
}
