package user

import "voice-app/domain"

type Service interface {
	GetAll() ([]domain.User, error)
	GetByPhoneNumber(phoneNumber string) (*domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service { return &service{repository: repository} }

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetByPhoneNumber(phoneNumber string) (*domain.User, error) {
	user, err := s.repository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}
	return user, nil
}
