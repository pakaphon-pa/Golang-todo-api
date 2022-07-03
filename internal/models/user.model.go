package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"unique" `
	Password  string    `json:"-"`
	Roles     []Role    `json:"roles" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

type UserServiceInterface interface {
	Get() ([]*User, error)
	GetById() (*User, error)
}

type UserRepositoryInterface interface {
	Find() ([]*User, error)
}
