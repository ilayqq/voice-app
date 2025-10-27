package domain

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"uniqueIndex"`
	FullName    string
	PhoneNumber string `gorm:"uniqueIndex"`
	Password    string `gorm:"not null"`
	Roles       []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"uniqueIndex;not null"`
	Users []*User `gorm:"many2many:user_roles;"`
}
