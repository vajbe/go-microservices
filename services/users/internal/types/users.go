package types

import "time"

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}
