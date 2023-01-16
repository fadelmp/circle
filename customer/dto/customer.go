package dto

// Database Design
type Customer struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
	Company Company `json:"company"`
	Base
}
