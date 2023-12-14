package types

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
