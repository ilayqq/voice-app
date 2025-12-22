package domain

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	FullName    string `gorm:"not null"`
	PhoneNumber string
	Password    string `gorm:"not null"`
	Roles       []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_roles;"`
}
