package models

import (
	"gorm.io/gorm"
)

type Pembelian struct {
	gorm.Model
	FakturID     uint            `json:"id_faktur"`
	Faktur       FakturPembelian `json:"faktur"`
	Quantity     uint            `json:"quantity"`
	TipeQuantity string          `json:"tipe_quantity"`
	Harga        uint            `json:"harga"`
	HargaPokok   uint            `json:"harga_pokok"`
	StockID      uint            `json:"stock_id"`
	Stock        Stock           `json:"stock"`
}

type Penjualan struct {
	gorm.Model
	FakturID     uint            `json:"id_faktur"`
	Faktur       FakturPenjualan `json:"faktur"`
	Quantity     uint            `json:"quantity"`
	TipeQuantity string          `json:"tipe_quantity"`
	Harga        uint            `json:"harga"`
	HargaPokok   uint            `json:"harga_pokok"`
	StockID      uint            `json:"stock_id"`
	Stock        Stock           `json:"stock"`
}
