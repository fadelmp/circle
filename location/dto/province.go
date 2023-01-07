package dto

// Database Design
type Province struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CountryID   uint   `json:"country_id"`
	CountryName string `json:"country_name"`
}
