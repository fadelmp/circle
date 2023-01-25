package entity

// Database Design
type Customer struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
	Phone      string `gorm:"type:VARCHAR(255);NOT NULL"`
	OtherPhone string `gorm:"type:VARCHAR(255)"`
	Email      string `gorm:"type:VARCHAR(255)"`
	Address    Address
	Base       `gorm:"embedded"`
}
