package entity

// Database Design
type City struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	ProvinceID uint   `gorm:"type:INT;NOT NULL;Index"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
	Province   Province
	Districts  []District
}
