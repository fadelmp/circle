package entity

// Database Design
type Order struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	Number     string `gorm:"type:VARCHAR(255);NOT NULL"`
	CustomerID uint   `gorm:"type:INT;NOT NULL;INDEX"`
	StatusID   uint   `gorm:"type:INT;NOT NULL;INDEX"`
	Total      int    `gorm:"type:INT;NOT NULL;default:0"`
	Note       string `gorm:"type:VARCHAR(255)"`
	Articles   []Article
	Base       `gorm:"embedded"`
}
