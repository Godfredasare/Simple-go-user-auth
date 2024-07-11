package model

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" validate:"required"`
	Currency    string   `json:"currency" validate:"required"`
	Quantity    int64    `json:"quantity"`
	Active      bool     `json:"active"`
	User_id     int64    `json:"user_id,omitempty"`
	Created_at  string   `json:"created_at" validate:"required"`
	Category    Category `json:"category"`
}
