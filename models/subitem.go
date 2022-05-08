package models

import (
	"gorm.io/gorm"
)

type Subitem struct {
	gorm.Model
	Int1    int    `json:"int1"`
	String1 string `json:"string1"`
	OtherID int    `json:"id_other"`
	Other   Item   `json:"other"`
}
