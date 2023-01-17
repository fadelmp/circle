package dto

// Database Design
type Status struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"front_description"`
	Base
}
