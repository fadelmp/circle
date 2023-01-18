package dto

// Database Design
type Service struct {
	ID        uint `json:"id"`
	ArticleID uint `json:"article_id"`
	ServiceID uint `json:"service_id"`
	Amount    int  `json:"amount"`
	Status    bool `json:"status"`
}
