package entity

// Database Design
type Address struct {
	ID            uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	CustomerID    uint   `gorm:"type:INT;NOT NULL;Index"`
	Line          string `gorm:"type:VARCHAR(255);"`
	SubDistrictID uint   `gorm:"type:INT;NOT NULL;Index"`
	DistrictID    uint   `gorm:"type:INT;NOT NULL;Index"`
	CityID        uint   `gorm:"type:INT;NOT NULL;Index"`
	ProvinceID    uint   `gorm:"type:INT;NOT NULL;Index"`
	CountryID     uint   `gorm:"type:INT;NOT NULL;Index"`
	PostalCode    string `gorm:"type:VARCHAR(255)"`
}
