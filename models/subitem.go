package models

import (
	"github.com/kamil5b/backend-template/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Subitem struct {
	gorm.Model
	Int1    int    `json:"int1"`
	String1 string `json:"string1"`
	OtherID int    `json:"id_other"`
	Other   Item   `json:"other"`
}

//POST SUBITEM

//GET AN ITEM

func GetASubitem(query string, val ...interface{}) (Subitem, error) {
	var subitem Subitem
	db := database.DB.Where(query, val).Preload(clause.Associations).First(&subitem)
	if db.Error != nil {
		return subitem, db.Error
	}
	return subitem, nil
}

//ARRAY

func GetPartSubitems(query string, val ...interface{}) ([]Subitem, error) {
	var subitems []Subitem
	db := database.DB.Where(query, val).Preload(clause.Associations).Find(&subitems)
	if db.Error != nil {
		return subitems, db.Error
	}
	return subitems, nil
}

//ALL ITEMS

func GetAllSubitems() ([]Subitem, error) {
	var subitems []Subitem
	db := database.DB.Preload(clause.Associations).Find(&subitems)
	if db.Error != nil {
		return subitems, db.Error
	}
	return subitems, nil
}
