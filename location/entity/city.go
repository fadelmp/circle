package entity

// Database Design
type City struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	ProvinceID uint   `gorm:"type:INT;NOT NULL;"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
	Districts  []District
}
