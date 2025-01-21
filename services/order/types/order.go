package types

// Order represents the orders table
type Product struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Order struct {
	ID            int       `json:"id"`             // Unique identifier for the order
	UserID        int       `json:"user_id"`        // User who placed the order
	OrderStatus   string    `json:"order_status"`   // Status of the order (PENDING, CONFIRMED, etc.)
	TotalAmount   float64   `json:"total_amount"`   // Total cost of the order
	PaymentStatus string    `json:"payment_status"` // Status of the payment (PENDING, PAID, etc.)
	CreatedAt     int64     `json:"created_at"`     // Order creation timestamp
	UpdatedAt     int64     `json:"updated_at"`     // Last updated timestamp
	Products      []Product `json:"products"`       // Store product details directly
}

type OrdertListingResponse struct {
	Result  []Order `json:"result"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	OrderBy string  `json:"order_by"`
	SortBy  string  `json:"sort_by"`
}
