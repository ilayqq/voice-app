package user

import (
	"voice-app/config"
	"voice-app/domain"
)

type Repository interface {
	Create(user *domain.User) error
}
type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(user *domain.User) error {
	return config.DB.Create(user).Error
}
