package dto

// Database Design
type Address struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	Line       string `json:"line"`
	DistrictID uint   `json:"district_id"`
	CityID     uint   `json:"city_id"`
	ProvinceID uint   `json:"province_id"`
	CountryID  uint   `json:"country_id"`
	PostalCode string `json:"postal_code"`
}
