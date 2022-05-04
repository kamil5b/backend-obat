package repositories

import (
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateItem(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := utilities.GoDotEnvVariable("LOG")
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.Item{
		Name:        data["name"],
		Description: data["description"],
	})
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	/*
		belum register. Register masukin ke table user hash
	*/
	msg = data["name"] + " berhasil terbuat"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====DELETE=====
func DeleteItem(data map[string]string, dataint map[string]int, IP string) error { //DELETE

	log := utilities.GoDotEnvVariable("LOG")
	msg := data["id"] + " menghapus"
	utilities.WriteLog(log, IP, msg)
	item, err := GetAnItem("ID = ?", dataint["id"])
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&item)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = data["id"] + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateItem(data map[string]string, dataint map[string]int, IP string) error { //DELETE

	log := utilities.GoDotEnvVariable("LOG")
	msg := data["id"] + " update"
	utilities.WriteLog(log, IP, msg)
	item, err := GetAnItem("ID = ?", dataint["id"])
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&item).Updates(models.Item{
		Name:        data["name"],
		Description: data["description"],
	})
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = data["id"] + " berhasil update"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//======GET========

//IS EXIST

func IsItemExist(query string, val ...interface{}) bool {
	var item models.Item
	database.DB.Where(query, val).Last(&item)
	return item.ID == 0
}

//GET AN ITEM

func GetAnItem(query string, val ...interface{}) (models.Item, error) {
	var item models.Item
	db := database.DB.Where(query, val).Last(&item)
	if db.Error != nil {
		return item, db.Error
	}
	return item, nil
}

//ARRAY

func GetPartItems(query string, val ...interface{}) ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Where(query, val).Find(&items)
	if db.Error != nil {
		return nil, db.Error
	}
	return items, nil
}

//ALL ITEMS

func GetAllItems() ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Find(&items)
	if db.Error != nil {
		return nil, db.Error
	}
	return items, nil
}
