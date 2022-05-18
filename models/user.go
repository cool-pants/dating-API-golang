package models

type User struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Location float64 `json:"location"`
	Gender   string  `json:"gender"`
	Email    string  `json:"email"`
}

type MatchList struct {
	UserID  int64  `json:"user_id"`
	Matches []User `json:"matches"`
}