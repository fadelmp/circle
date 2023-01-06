package entity

// Database Design
type Country struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name      string `gorm:"type:VARCHAR(255);NOT NULL"`
	Provinces []Province
}
