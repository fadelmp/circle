package dto

// Database Design
type Customer struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	OtherPhone string  `json:"other_phone"`
	Email      string  `json:"email"`
	Address    Address `json:"address"`
	Base
}

type ShowCustomer struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	OtherPhone  string `json:"other_phone"`
	Email       string `json:"email"`
	AddressLine string `json:"address_line"`
	IsActived   bool   `json:"is_actived"`
}
