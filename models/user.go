package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}
