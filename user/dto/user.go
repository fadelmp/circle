package dto

// Database Design
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint
	Base     `gorm:"embedded"`
}
