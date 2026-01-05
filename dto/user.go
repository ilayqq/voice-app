package dto

type UserResponse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	FullName    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique;not null"`
}
