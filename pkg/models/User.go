package models

type User struct {
	ID       string `json:"id" db:"id,omitempty"`
	Username string `json:"username" db:"username"`
}
