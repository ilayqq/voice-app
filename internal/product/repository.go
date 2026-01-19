package product

import (
	"voice-app/config"
	"voice-app/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetByBarcode(code string) (*domain.Product, error)
	Create(product *domain.Product) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	config.DB.Preload("Stocks").Find(&products)
	return products, nil
}

func (r *repository) GetByBarcode(code string) (*domain.Product, error) {
	var product domain.Product
	result := config.DB.Where("barcode = ?", code).First(&product)
	return &product, result.Error
}

func (r *repository) Create(product *domain.Product) error {
	return config.DB.Create(&product).Error
}
