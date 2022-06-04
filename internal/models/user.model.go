package models

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"title"`
}

type UserServiceInterface interface {
	Get() ([]*User, error)
	GetById() (*User, error)
}

type UserRepositoryInterface interface {
	Find() ([]*User, error)
}
