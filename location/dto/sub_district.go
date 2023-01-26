package dto

// Database Design
type SubDistrict struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	DistrictID   uint   `json:"district_id"`
	DistrictName string `json:"district_name"`
}
