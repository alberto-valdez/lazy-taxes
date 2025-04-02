package services

import (
	"time"

	"github.com/alberto-valdez/lazy-taxes/internal/models"
)

type UserService interface {
	GetUser() (string, error)
	CreateUser(user models.User) (models.User, error)
}

type userService struct{}

func NewItemService() UserService {
	return &userService{}
}

func (s *userService) GetUser() (string, error) {
	return "init", nil
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return models.User{
		ID:          "",
		Username:    user.Username,
		Email:       user.Email,
		ActiveLogin: "",
		FullName:    "",
		Role:        "",
		Active:      false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}, nil
}
