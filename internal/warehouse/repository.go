package warehouse

import (
	"voice-app/config"
	"voice-app/domain"
)

type Repository interface {
	GetAll() ([]domain.Warehouse, error)
	Create(warehouse *domain.Warehouse) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	result := config.DB.Preload("Owner").Preload("Stocks").Find(&warehouses)
	if result.Error != nil {
		return nil, result.Error
	}

	return warehouses, nil
}

func (r *repository) Create(warehouse *domain.Warehouse) error {
	return config.DB.Create(&warehouse).Error
}
