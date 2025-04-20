package services

import (
	"errors"

	"github.com/yagoayala/server/internal/models"
	"github.com/yagoayala/server/internal/repository"
)

type UserService interface {
	CreateUser(*models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(uint) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	u, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.userRepo.DeleteUser(u)
}
