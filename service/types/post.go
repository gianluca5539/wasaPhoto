package types

// Post represents a post.
type Post struct {
	PostID      int64  `json:"postid"`
	UserID      int    `json:"userid"`
	Username    string `json:"username"`
	Feeling     int    `json:"feeling"`
	UserPicture int    `json:"userpicture"`
	Picture     int    `json:"picture"`
	Caption     string `json:"caption"`
	CreatedAt   int    `json:"createdat"`
	LikeCount   int    `json:"likecount"`
}
