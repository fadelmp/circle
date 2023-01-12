package entity

import "time"

// Database Design
type Customer struct {
	ID             uint      `gorm:"primaryKey;autoIncrement:true;Index"`
	Name           string    `gorm:"type:VARCHAR(255);NOT NULL"`
	Phone          string    `gorm:"type:VARCHAR(255);NOT NULL"`
	Email          string    `gorm:"type:VARCHAR(255)"`
	IsActive       bool      `gorm:"type:TINYINT(1);default:true"`
	CreatedBy      string    `gorm:"type:VARCHAR(255);NOT NULL"`
	ModifyBy       string    `gorm:"type:VARCHAR(255);NOT NULL"`
	CreatedAt      time.Time `gorm:"type:DATETIME;NOT NULL;default:now()"`
	ModifyAt       time.Time `gorm:"type:DATETIME;NOT NULL;default:now()"`
	Addresses      []Address
	ContactPeoples []ContactPeople
}
