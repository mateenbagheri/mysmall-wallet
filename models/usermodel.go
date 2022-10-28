package models

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type UserBalance struct {
	Balance float64 `json:"balance"`
}
