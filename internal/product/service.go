package product

import "voice-app/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Create(product domain.Product) (domain.Product, error) {
	if err := s.repository.Create(&product); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
