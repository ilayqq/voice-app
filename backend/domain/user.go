package domain

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	FullName    string `gorm:"not null" json:"fullName"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `gorm:"not null"`
	Roles       []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_roles;"`
}
