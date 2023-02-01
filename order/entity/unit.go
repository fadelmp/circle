package entity

// Database Design
type Unit struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"type:VARCHAR(50);NOT NULL"`
	Description string `gorm:"type:VARCHAR(255)"`
}
