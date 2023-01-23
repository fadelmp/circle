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

type ShowCustomer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	CompanyName  string `json:"company_name"`
	CompanyPhone string `json:"company_phone"`
	AddressLine  string `json:"address_line"`
	IsActived    bool   `json:"is_actived"`
}
