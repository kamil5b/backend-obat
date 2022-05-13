package models

import (
	"time"

	"gorm.io/gorm"
)

type Piutang struct {
	gorm.Model
	FakturID   uint            `json:"id_faktur"`
	Faktur     FakturPenjualan `json:"faktur"`
	CustomerID uint            `json:"id_customer"`
	Customer   Customer        `json:"customer"`
	Nominal    int             `json:"nominal"`
	Tanggal    time.Time       `json:"tanggal"`
	JatuhTempo time.Time       `json:"jatuh_tempo"`
}

type SummaryPiutang struct {
	gorm.Model
	FakturID       uint            `json:"id_faktur"`
	Faktur         FakturPenjualan `json:"faktur"`
	CustomerID     uint            `json:"id_customer"`
	Customer       Customer        `json:"customer"`
	TotalNominal   int             `json:"total_nominal"`
	TanggalPiutang time.Time       `json:"tanggal_piutang"`
	JatuhTempo     time.Time       `json:"jatuh_tempo"`
}
