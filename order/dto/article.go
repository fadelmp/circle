package dto

// Database Design
type Article struct {
	ID         uint   `json:"id"`
	OrderID    uint   `json:"order_id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	UnitID     uint   `json:"unit_id"`
	UnitName   string `json:"unit_name"`
	Amount     int    `json:"amount"`
	Note       string `json:"note"`
	Path_Image string `json:"path_image"`
	StatusID   uint   `json:"status_id"`
	StatusName string `json:"status_name"`
	Services   []Service
}
