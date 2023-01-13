package entity

// Database Design
type Service struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name        string `gorm:"type:VARCHAR(255);NOT NULL"`
	Description string `gorm:"type:VARCHAR(255)"`
	Price       string `gorm:"type:INT; NOT NULL; default: 0"`
	Base        `gorm:"embedded"`
}
