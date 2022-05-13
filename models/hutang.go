package models

import (
	"time"

	"gorm.io/gorm"
)

type Hutang struct {
	gorm.Model
	FakturID   uint            `json:"id_faktur"`
	Faktur     FakturPembelian `json:"faktur"`
	SuplierID  uint            `json:"id_suplier"`
	Suplier    Suplier         `json:"suplier"`
	Nominal    int             `json:"nominal"`
	Tanggal    time.Time       `json:"tanggal"`
	JatuhTempo time.Time       `json:"jatuh_tempo"`
}

type SummaryHutang struct {
	gorm.Model
	FakturID      uint            `json:"id_faktur"`
	Faktur        FakturPembelian `json:"faktur"`
	SuplierID     uint            `json:"id_suplier"`
	Suplier       Suplier         `json:"suplier"`
	TotalNominal  int             `json:"total_nominal"`
	TanggalHutang time.Time       `json:"tanggal_hutang"`
	JatuhTempo    time.Time       `json:"jatuh_tempo"`
}
