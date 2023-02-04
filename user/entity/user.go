package entity

// Database Design
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Name     string `gorm:"type:VARCHAR(255);NOT NULL"`
	Username string `gorm:"type:VARCHAR(255);NOT NULL"`
	Password string `gorm:"type:VARCHAR(255)"`
	RoleID   uint
	Base     `gorm:"embedded"`
	Role     Role
}
