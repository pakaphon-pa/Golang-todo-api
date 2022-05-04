package repository

import (
	"gotaskapp/internal/models"
)

type userRepository struct {
}

func (u *userRepository) Find() ([]*models.User, error) {

	return make([]*models.User, 0), nil
}

func NewUserRepository() models.UserRepositoryInterface {
	return &userRepository{}
}
