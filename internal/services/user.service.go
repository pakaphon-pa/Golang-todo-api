package services

import (
	"fmt"
	"gotaskapp/internal/models"
)

type userService struct {
	repo models.UserRepositoryInterface
}

func (s *userService) Get() ([]*models.User, error) {
	fmt.Println("1111111111111111111")
	return s.repo.Find()
}

func NewUserService(userRepo models.UserRepositoryInterface) models.UserServiceInterface {
	return &userService{
		repo: userRepo,
	}
}
