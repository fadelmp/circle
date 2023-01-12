package dto

import "time"

// Database Design
type ContactPeople struct {
	ID         uint      `json:"id"`
	CustomerID uint      `json:"customer_id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	IsActive   bool      `json:"is_active"`
	CreatedBy  string    `json:"created_by"`
	ModifyBy   string    `json:"modify_by"`
	CreatedAt  time.Time `json:"created_at"`
	ModifyAt   time.Time `json:"modify_at"`
}
