package models

import (
	"errors"
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/utilities"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

//POST

func CreateItem(data map[string]string, dataint map[string]int, IP string) error { //POST

	var item Item
	log := utilities.GoDotEnvVariable("LOG")
	msg := data["username"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)

	db := database.DB.Where("ID = ?", dataint["id"]).First(&item)
	if db.Error != nil {
		msg = data["name"] + " " + db.Error.Error()
		utilities.WriteLog(log, IP, msg)
		return db.Error

	} else if item.ID == 0 {
		//CREATE FORM
		database.DB.Create(&Item{
			Name:        data["name"],
			Description: data["description"],
		})
		msg = data["id_student"] + " berhasil form"
		utilities.WriteLog(log, IP, msg)
		return nil
	}

	/*
		belum register. Register masukin ke table user hash
	*/
	msg = strconv.Itoa(dataint["id"]) + " telah di input sebelumnya"
	utilities.WriteLog(log, IP, msg)
	return errors.New(msg)
}

//======GET========

//IS EXIST

func IsItemExist(query string, val ...interface{}) bool {
	var item Item
	database.DB.Where(query, val).First(&item)
	return item.ID == 0
}

//GET AN ITEM

func GetAnItem(query string, val ...interface{}) (Item, error) {
	var item Item
	db := database.DB.Where(query, val).First(&item)
	if db.Error != nil {
		return item, db.Error
	}
	return item, nil
}

//ARRAY

func GetPartItems(query string, val ...interface{}) ([]Item, error) {
	var items []Item
	db := database.DB.Where(query, val).Find(&items)
	if db.Error != nil {
		return nil, db.Error
	}
	return items, nil
}

//ALL ITEMS

func GetAllItems() ([]Item, error) {
	var items []Item
	db := database.DB.Find(&items)
	if db.Error != nil {
		return nil, db.Error
	}
	return items, nil
}
