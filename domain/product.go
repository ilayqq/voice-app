package domain

import "time"

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	SKU         string `json:"SKU" gorm:"unique;not null"`
	Description string
	Category    string
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	Stocks []Stock `json:"stocks" gorm:"foreignKey:ProductID"`
}
