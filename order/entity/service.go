package entity

// Database Design
type Service struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	ArticleID uint `gorm:"type:INT;NOT NULL;INDEX"`
	ServiceID uint `gorm:"type:INT;NOT NULL;INDEX"`
	Amount    int  `gorm:"type:INT;NOT NULL;INDEX"`
	Status    bool `gorm:"type:tinyint(1);NOT NULL;default:true"`
}
