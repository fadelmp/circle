package dto

// Database Design
type Order struct {
	ID         uint   `json:"id"`
	Number     string `json:"order_number"`
	CustomerID uint   `json:"customer_id"`
	StatusID   uint   `json:"status_id"`
	Amount     int    `json:"amount"`
	Note       string `json:"note"`
	Base
	Articles []Article
}
