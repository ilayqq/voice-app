package auth

import (
	"errors"
	"os"
	"time"
	"voice-app/domain"
	"voice-app/internal/user"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(phoneNumber, password string, roles []string) (*domain.User, error)
	Login(phoneNumber, password string) (string, error)
}
type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Register(phoneNumber, password string, roles []string) (*domain.User, error) {
	existingUser, err := s.repository.Exist(phoneNumber)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user with this phone number already exists")
	}

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

func (s *service) Login(phoneNumber, password string) (string, error) {
	user, err := s.repository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Name
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         user.ID,
		"phoneNumber": user.PhoneNumber,
		"roles":       roleNames,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
