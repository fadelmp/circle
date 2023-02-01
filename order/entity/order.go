package entity

import "time"

// Database Design
type Order struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true"`
	Number        string    `gorm:"type:VARCHAR(255);NOT NULL"`
	CustomerID    uint      `gorm:"type:INT;NOT NULL;INDEX"`
	CustomerName  string    `gorm:"type:VARCHAR(255)"`
	StatusID      uint      `gorm:"type:INT;NOT NULL;INDEX;default:1"`
	Amount        int       `gorm:"type:INT;NOT NULL;default:0"`
	Type          string    `gorm:"type:VARCHAR(255)"`
	Date          time.Time `gorm:"type:datetime;default:now()"`
	DeliveryOrder string    `gorm:"type:VARCHAR(255)"`
	Note          string    `gorm:"type:VARCHAR(255)"`
	Base          `gorm:"embedded"`
	Articles      []Article
	Status        Status
}
