<<<<<<< HEAD
package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateItem(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
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
func DeleteItem(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	item, err := GetItem("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&item)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateItem(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	item, err := GetItem("ID = ?", ID)
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

	msg = strconv.Itoa(ID) + " berhasil update"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//======GET========

//IS EXIST

func IsItemExist(query string, val ...interface{}) bool {
	var item models.Item
	database.DB.Where(query, val...).Last(&item)
	return item.ID != 0
}

//GET AN ITEM

func GetItem(query string, val ...interface{}) (models.Item, error) {
	var item models.Item
	db := database.DB.Where(query, val...).Last(&item)
	if db.Error != nil {
		return item, db.Error
	}
	return item, nil
}

//ARRAY

func GetItems(query string, val ...interface{}) ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Where(query, val...).Find(&items)
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
=======
package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateItem(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
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
func DeleteItem(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	item, err := GetItem("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&item)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateItem(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	item, err := GetItem("ID = ?", ID)
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

	msg = strconv.Itoa(ID) + " berhasil update"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//======GET========

//IS EXIST

func IsItemExist(query string, val ...interface{}) bool {
	var item models.Item
	database.DB.Where(query, val...).Last(&item)
	return item.ID != 0
}

//GET AN ITEM

func GetItem(query string, val ...interface{}) (models.Item, error) {
	var item models.Item
	db := database.DB.Where(query, val...).Last(&item)
	if db.Error != nil {
		return item, db.Error
	}
	return item, nil
}

//ARRAY

func GetItems(query string, val ...interface{}) ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Where(query, val...).Find(&items)
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
>>>>>>> 3b907fbf44e0be985c180d88d513e8a91064017c
