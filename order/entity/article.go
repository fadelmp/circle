package entity

// Database Design
type Article struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	OrderID    uint   `gorm:"type:INT;NOT NULL;INDEX"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL;INDEX"`
	Quantity   int    `gorm:"type:INT;NOT NULL"`
	Unit       string `gorm:"type:VARCHAR(255)"`
	Amount     int    `gorm:"type:INT;NOT NULL"`
	Note       string `gorm:"type:VARCHAR(255)"`
	Path_Image string `gorm:"type:VARCHAR(255)"`
	Status     bool   `gorm:"type:tinyint(1);NOT NULL;default:true"`
	Services   []Service
}
