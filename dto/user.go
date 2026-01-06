package dto

type UserResponse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	PhoneNumber string `gorm:"unique;not null" json:"phone_number"`

	RoleName string `json:"role"`
}
