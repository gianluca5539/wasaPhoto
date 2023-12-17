package types

type UserProfile struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Picture   int        `json:"picture"`
	Feeling   int        `json:"feeling"`
	Bio       string     `json:"bio"`
	Followers []User     `json:"followers"`
	Following []User     `json:"following"`
	Banned    bool       `json:"banned"`
	Posts     []UserPost `json:"posts"`
}
