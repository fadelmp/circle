package entity

// Database Design
type ArticleService struct {
	ID             uint `gorm:"primaryKey;autoIncrement:true"`
	OrderArticleID uint `gorm:"type:INT;NOT NULL;INDEX"`
	ServiceID      uint `gorm:"type:INT;NOT NULL;INDEX"`
	Price          int  `gorm:"type:INT;NOT NULL;INDEX"`
	Status         bool `gorm:"type:tinyint(1);NOT NULL;default:true"`
	Base           `gorm:"embedded"`
}
