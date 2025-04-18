package services

import (
	"errors"
	"goda/pkg/models"
	"goda/pkg/repositories"
)

type UserService interface {
	Register(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) Register(user *models.User) error {
	// Example: unique email validation (you can extend validations like this)
	existingUser, _ := s.repository.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}
	return s.repository.Create(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repository.FindByID(id)
}
