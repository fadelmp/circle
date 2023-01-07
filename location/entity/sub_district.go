package entity

// Database Design
type SubDistrict struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	DistrictID uint   `gorm:"type:INT;NOT NULL;Index"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
}
