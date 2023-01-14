package dto

// Database Design
type Service struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"service_code"`
	Price       int    `json:"price"`
	Base
}
