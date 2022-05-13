package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	BarangID   uint      `json:"kode_barang"`
	Barang     Barang    `json:"barang"`
	Expired    time.Time `json:"expired"`
	BigQty     uint      `json:"big_quantity"`
	MediumQty  uint      `json:"medium_quantity"`
	SmallQty   uint      `json:"small_quantity"`
	HargaPokok uint      `json:"harga_pokok"`
}
