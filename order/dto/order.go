package dto

import "time"

// Database Design
type Order struct {
	ID           uint   `json:"id"`
	Number       string `json:"order_number"`
	CustomerID   uint   `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	StatusID     uint   `json:"status_id"`
	Amount       int    `json:"amount"`
	Note         string `json:"note"`
	OrderType    string `json:"type"`
	OrderDate    string `json:"date"`
	Articles     []Article
	Base
}

type ShowOrder struct {
	ID           uint      `json:"id"`
	Number       string    `json:"number"`
	CustomerName string    `json:"customer_name"`
	StatusName   string    `json:"status_name"`
	ArticleCount int       `json:"article_count"`
	Amount       int       `json:"amount"`
	Note         string    `json:"note"`
	OrderDate    time.Time `json:"order_date"`
	OrderBy      string    `json:"order_by"`
	Articles     []Article
}
