package product

import (
	"voice-app/config"
	"voice-app/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Create(product *domain.Product) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	config.DB.Find(&products)
	return products, nil
}

func (r *repository) Create(product *domain.Product) error {
	return config.DB.Create(&product).Error
}
