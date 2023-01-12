package dto

import "time"

// Database Design
type Customer struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	Phone          string          `json:"phone"`
	Email          string          `json:"email"`
	IsActive       bool            `json:"is_active"`
	CreatedBy      string          `json:"created_by"`
	ModifyBy       string          `json:"modify_by"`
	CreatedAt      time.Time       `json:"created_at"`
	ModifyAt       time.Time       `json:"modify_at"`
	Addresses      []Address       `json:"addresses"`
	ContactPeoples []ContactPeople `json:"contact_peoples"`
}

type ShowCustomer struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	AddressLine   string `json:"address_line"`
	ContactPeople string `json:"contact_people"`
	IsActive      bool   `json:"is_active"`
}