package entity

// Database Design
type Company struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	CustomerID uint   `gorm:"type:INT;NOT NULL;Index"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
	Phone      string `gorm:"type:VARCHAR(255)"`
}
