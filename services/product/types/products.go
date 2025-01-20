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

type ProductListingResponse struct {
	Result  []ProductResponse `json:"result"`
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	OrderBy string            `json:"order_by"`
	SortBy  string            `json:"sort_by"`
}
