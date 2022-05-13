package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `json:"name"`
	Alamat       string `json:"alamat"`
	NomorHP      string `json:"nomor_hp"`
	LimitPiutang uint   `json:"limit_piutang"`
}
