package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NIK      string `json:"nik" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Verified bool   `json:"-"`
}
