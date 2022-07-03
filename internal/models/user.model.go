package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	RoleId   uint   `json:"role_id"`
	Roles    []Role `json:"roles" gorm:"foreignKey: RoleId"`
}

type UserServiceInterface interface {
	Get() ([]*User, error)
	GetById() (*User, error)
}

type UserRepositoryInterface interface {
	Find() ([]*User, error)
}
