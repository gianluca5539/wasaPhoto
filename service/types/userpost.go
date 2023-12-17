package types

// UserPost represents a user post.
type UserPost struct {
	PostID    int    `json:"postid"`
	Caption   string `json:"caption"`
	Picture   int    `json:"picture"`
	CreatedAt int    `json:"createdat"`
	LikeCount int    `json:"likecount"`
}
