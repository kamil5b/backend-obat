package models

import (
	"time"

	"gorm.io/gorm"
)

type TargetSales struct {
	gorm.Model
	UserID        uint      `json:"id_user"`
	User          User      `json:"user"`
	NominalTarget uint      `json:"nominal_target"`
	TanggalTarget time.Time `json:"tanggal_target"`
}

type Sales struct {
	gorm.Model
	UserID          uint            `json:"id_user"`
	User            User            `json:"user"`
	FakturID        uint            `json:"id_faktur"`
	Faktur          FakturPenjualan `json:"faktur"`
	TotalPenjualan  uint            `json:"total_penjualan"`
	NominalInsentif uint            `json:"nominal_insentif"`
	TanggalTurun    time.Time       `json:"tanggal_turun"`
}
