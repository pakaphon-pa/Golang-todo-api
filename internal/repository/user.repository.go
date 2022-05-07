package repository

import (
	"gotaskapp/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Find() ([]*models.User, error) {
	var data []*models.User

	u.db.Find(&data)
	return data, nil
}

func NewUserRepository(db *gorm.DB) models.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
