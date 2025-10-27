package auth

import (
	"voice-app/domain"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(phoneNumber, password string, roles []string) (*domain.User, error)
}
type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) Register(phoneNumber, password string, roles []string) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		PhoneNumber: phoneNumber,
		Password:    string(hash),
		Roles:       make([]*domain.Role, len(roles)),
	}

	for i, roleName := range roles {
		user.Roles[i] = &domain.Role{Name: roleName}
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
}
