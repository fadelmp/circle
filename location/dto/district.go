package dto

// Database Design
type District struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	CityID   uint   `json:"city_id"`
	CityName string `json:"city_name"`
}
