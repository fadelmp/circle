package dto

// Database Design
type Company struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}
