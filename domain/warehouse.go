package domain

import "time"

type Warehouse struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Location    string
	Description string
	OwnerID     uint
	Owner       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Stocks []Stock `gorm:"foreignKey:WarehouseID"`
}
