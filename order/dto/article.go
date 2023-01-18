package dto

// Database Design
type Article struct {
	ID         uint   `json:"id"`
	OrderID    uint   `json:"order_id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Amount     int    `json:"amount"`
	Note       string `json:"note"`
	Path_Image string `json:"path_image"`
	Status     bool   `json:"status"`
	Services   []Service
}
