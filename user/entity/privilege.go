package entity

// Database Design
type Privilege struct {
	ID    uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name  string `gorm:"type:VARCHAR(255)"`
	Base  `gorm:"embedded"`
	Roles []Role `gorm:"many2many:role_privileges"`
}
