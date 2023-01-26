package entity

// Database Design
type District struct {
	ID           uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	CityID       uint   `gorm:"type:INT;NOT NULL;Index"`
	Name         string `gorm:"type:VARCHAR(255);NOT NULL"`
	City         City
	SubDistricts []SubDistrict
}
