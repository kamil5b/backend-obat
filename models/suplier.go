package models

import (
	"gorm.io/gorm"
)

type Suplier struct {
	gorm.Model
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	NomorHP string `json:"nomor_hp"`
}
