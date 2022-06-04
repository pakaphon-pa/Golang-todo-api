package services

import (
	"gotaskapp/internal/models"
	"gotaskapp/pkg/customError"
	"net/http"
)

type userService struct {
	repo models.UserRepositoryInterface
}

func (s *userService) Get() ([]*models.User, error) {

	return s.repo.Find()
}

func (s *userService) GetById() (*models.User, error) {
	return nil, customError.NewHTTPError(http.StatusBadRequest, "Invalid", "invalid user id")
}

func NewUserService(userRepo models.UserRepositoryInterface) models.UserServiceInterface {
	return &userService{
		repo: userRepo,
	}
}
