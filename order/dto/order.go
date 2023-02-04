package dto

import "time"

// Database Design
type Order struct {
	ID            uint      `json:"id"`
	Number        string    `json:"order_number"`
	CustomerID    uint      `json:"customer_id"`
	CustomerName  string    `json:"customer_name"`
	StatusID      uint      `json:"status_id"`
	Amount        int       `json:"amount"`
	Type          string    `json:"type"`
	Date          time.Time `json:"date"`
	DeliveryOrder string    `json:"delivery_order"`
	Note          string    `json:"note"`
	Articles      []Article
	Base
}

type ShowOrder struct {
	ID           uint      `json:"id"`
	Number       string    `json:"number"`
	CustomerName string    `json:"customer_name"`
	StatusName   string    `json:"status_name"`
	ArticleCount int       `json:"article_count"`
	Amount       int       `json:"amount"`
	Type         string    `json:"type"`
	Date         time.Time `json:"date"`
	OrderBy      string    `json:"order_by"`
}

type ShowOrderDetail struct {
	ID           uint      `json:"id"`
	Number       string    `json:"number"`
	CustomerID   uint      `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	StatusName   string    `json:"status_name"`
	ArticleCount int       `json:"article_count"`
	Amount       int       `json:"amount"`
	Type         string    `json:"type"`
	Date         time.Time `json:"date"`
	OrderBy      string    `json:"order_by"`
	Articles     []Article `json:"articles"`
}
