package models

import (
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Kode          string `json:"kode" gorm:"unique"`
	Name          string `json:"name"`
	TipeBig       string `json:"tipe_big"`
	TipeMedium    string `json:"tipe_medium"`
	TipeSmall     string `json:"tipe_small"`
	BigToMedium   uint   `json:"big_to_small"`
	MediumToSmall uint   `json:"medium_to_small"`
	HargaBig      string `json:"harga_big"`
	HargaMedium   string `json:"harga_medium"`
	HargaSmall    string `json:"harga_small"`
}
