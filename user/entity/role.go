package entity

// Database Design
type Role struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name       string `gorm:"type:VARCHAR(255);NOT NULL"`
	Base       `gorm:"embedded"`
	Privileges []Privilege `gorm:"many2many:role_privileges"`
}
