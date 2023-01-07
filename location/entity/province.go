package entity

// Database Design
type Province struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	CountryID uint   `gorm:"type:INT;NOT NULL;"`
	Name      string `gorm:"type:VARCHAR(255);NOT NULL"`
	Country   Country
	Cities    []City
}
