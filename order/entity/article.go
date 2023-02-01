package entity

// Database Design
type Article struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	OrderID    uint   `gorm:"type:INT;NOT NULL;INDEX"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL;INDEX"`
	Quantity   int    `gorm:"type:INT;NOT NULL"`
	UnitID     uint   `gorm:"type:INT"`
	Amount     int    `gorm:"type:INT;NOT NULL"`
	Note       string `gorm:"type:VARCHAR(255)"`
	Path_Image string `gorm:"type:VARCHAR(255)"`
	StatusID   uint   `gorm:"type:INT;NOT NULL;INDEX;default:1"`
	Services   []Service
	Status     Status
	Order      Order
	Unit       Unit
}
