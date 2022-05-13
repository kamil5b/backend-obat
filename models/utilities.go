package models

import (
	"time"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Nomor uint   `json:"nomor_bank"`
	Nama  string `json:"nama_bank"`
}

type Giro struct {
	gorm.Model
	Tanggal   time.Time `json:"tanggal"`
	NomorGiro string    `json:"nomor_giro"`
	BankID    uint      `json:"id_bank"`
	Bank      Bank      `json:"bank"`
	Nominal   uint      `json:"nominal"`
}

type KasKecil struct {
	gorm.Model
	Nominal     int `json:"nominal"`
	Description int `json:"description"`
}
