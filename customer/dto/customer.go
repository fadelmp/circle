package dto

// Database Design
type Customer struct {
	ID          uint    `json:"id" param:"id"`
	Name        string  `json:"name"`
	Phone       string  `json:"phone"`
	OtherPhone  string  `json:"other_phone"`
	Email       string  `json:"email"`
	AddressLine string  `json:"address_line"`
	IsActived   bool    `json:"is_actived" param:"is_actived"`
	Address     Address `json:"address"`
}
