package domain

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	PhoneNumber string `gorm:"unique;not null" json:"phone_number"`
	Password    string `gorm:"not null"`
	Roles       []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_roles;"`
}
