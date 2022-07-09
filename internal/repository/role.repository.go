package repository

import (
	"gotaskapp/internal/models"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) Create(role *models.Role) (*models.Role, error) {
	r.db.Create(&role)

	return role, nil
}

func (r *roleRepository) Get() ([]*models.Role, error) {
	var data []*models.Role
	r.db.Find(&data)
	return data, nil
}

func (r *roleRepository) GetById(id uint) (*models.Role, error) {
	var data *models.Role
	r.db.Model(data).Where("id = ?", id)
	return data, nil
}

func NewRoleRepository(db *gorm.DB) models.RoleRepositoryInterface {
	return &roleRepository{
		db: db,
	}
}
