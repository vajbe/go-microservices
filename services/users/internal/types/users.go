package types

import "time"

type User struct {
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Phone     string    `json:"phone_number"`
	Password  string    `json:"password_hash"`
	Id        int       `json:"id"`
}
