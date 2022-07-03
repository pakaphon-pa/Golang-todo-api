package models

type Role struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
