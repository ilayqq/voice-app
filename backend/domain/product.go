package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	SKU         string `gorm:"unique;not null"`
	Description string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Stocks []Stock `gorm:"foreignKey:ProductID"`
}
