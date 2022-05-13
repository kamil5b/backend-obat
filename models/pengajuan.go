package models

import (
	"gorm.io/gorm"
)

type Demand struct {
	gorm.Model
	BarangID          string `json:"kode"`
	Barang            Barang `json:"barang"`
	QuantityDemand    uint   `json:"quantity_demand"`
	QuantityAvailable uint   `json:"quantity_available"`
	TipeQuantity      string `json:"tipe_quantity"`
}

type Pengajuan struct {
	gorm.Model
	BarangID      string `json:"kode"`
	Barang        Barang `json:"barang"`
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
