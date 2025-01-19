package types

type Product struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	CreatedAt   int64   `json:"created_at"`
}

// ProductResponse struct for responses
type ProductResponse struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}
