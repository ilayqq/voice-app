package auth

import (
	"voice-app/domain"
	"voice-app/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(phoneNumber, password string, roles []string) (*domain.User, error)
}
type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Register(phoneNumber, password string, roles []string) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		PhoneNumber: phoneNumber,
		Password:    string(hash),
		Roles:       make([]domain.Role, len(roles)),
	}

	for i, roleName := range roles {
		user.Roles[i] = domain.Role{Name: roleName}
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
