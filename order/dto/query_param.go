package dto

import "time"

type QueryParam struct {
	CustomerID string    `query:"customer_id"`
	StatusID   string    `query:"status_id"`
	Search     string    `query:"search"`
	From       time.Time `query:"from"`
	To         time.Time `query:"to"`
}
