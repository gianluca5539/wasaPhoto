package types

// User represents a user.
type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Feeling  int    `json:"feeling"`
	Bio      string `json:"bio"`
	Picture  int    `json:"picture"`
}
