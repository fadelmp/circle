package dto

// Database Design
type Unit struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActived   bool   `json:"is_actived" param:"is_actived"`
}
