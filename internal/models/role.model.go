package models

type Role struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"many2many:user_roles"`
}

type RoleRequest struct {
	Name string `json:"name" validate:"required,len=3"`
}

type RoleServiceInterface interface {
	Find() ([]*Role, error)
	Save() (*Role, error)
}

type RoleRepositoryInterface interface {
	Get() ([]*Role, error)
	GetById(id uint) (*Role, error)
	Create(role *Role) (*Role, error)
}
