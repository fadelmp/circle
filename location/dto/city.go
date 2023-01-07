package dto

// Database Design
type City struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProvinceID   uint   `json:"province_id"`
	ProvinceName string `json:"province_name"`
}
