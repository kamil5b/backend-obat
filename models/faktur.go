package models

import (
	"time"

	"gorm.io/gorm"
)

type FakturPenjualan struct {
	gorm.Model
	Tanggal        time.Time `json:"tanggal"`
	TipePembayaran string    `json:"tipe_pembayaran"`
	CustomerID     uint      `json:"id_customer"`
	Customer       Customer  `json:"customer"`
	GiroID         uint      `json:"nomor_giro"`
	Giro           Giro      `json:"giro"`
	JatuhTempo     time.Time `json:"jatuh_tempo"`
}

type FakturPembelian struct {
	gorm.Model
	NomorFaktur    uint      `json:"nomor_faktur" gorm:"unique"`
	Tanggal        time.Time `json:"tanggal"`
	TipePembayaran string    `json:"tipe_pembayaran"`
	SuplierID      uint      `json:"id_suplier"`
	Suplier        Suplier   `json:"suplier"`
	GiroID         uint      `json:"nomor_giro"`
	Giro           Giro      `json:"giro"`
	JatuhTempo     time.Time `json:"jatuh_tempo"`
}
