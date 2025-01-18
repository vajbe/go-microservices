package types

// User struct with validation tags
type User struct {
	Name      string `json:"name,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Phone     string `json:"phone_number" validate:"required"`
	Password  string `json:"password,omitempty" validate:"required,min=8"`
	Id        string `json:"id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

// UserResponse struct for responses
type UserResponse struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone_number"`
	Id    string `json:"id,omitempty"`
}
