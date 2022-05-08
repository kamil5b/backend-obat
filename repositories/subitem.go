package repositories

import (
	"errors"
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
	"gorm.io/gorm/clause"
)

//=====POST======

func CreateSubitem(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := strconv.Itoa(dataint["int1"]) + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	item, err := GetItem("ID = ?", dataint["id_other"])
	if err != nil {
		msg = data["name"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return err

	} else if item.ID == 0 {
		//CREATE FORM
		msg = strconv.Itoa(dataint["id_other"]) + " tidak ditemukan"
		utilities.WriteLog(log, IP, msg)
		return errors.New(msg)
	}

	/*
		belum register. Register masukin ke table user hash
	*/
	database.DB.Create(&models.Subitem{
		Int1:    dataint["int1"],
		String1: data["string1"],
		OtherID: dataint["id_other"],
		Other:   item,
	})
	msg = data["name"] + " berhasil dibuat"
	utilities.WriteLog(log, IP, msg)
	return nil

}

//=====DELETE=====
func DeleteSubitem(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	subitem, err := GetSubitem("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&subitem)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateSubitem(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	subitem, err := GetSubitem("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	item, err := GetItem("ID = ?", dataint["id_other"])
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&subitem).Updates(models.Subitem{
		Int1:    dataint["int1"],
		String1: data["string1"],
		OtherID: dataint["id_other"],
		Other:   item,
	})
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil update"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//======GET======
//GET AN ITEM

func GetSubitem(query string, val ...interface{}) (models.Subitem, error) {
	var subitem models.Subitem
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&subitem)
	if db.Error != nil {
		return subitem, db.Error
	}
	return subitem, nil
}

//ARRAY

func GetSubitems(query string, val ...interface{}) ([]models.Subitem, error) {
	var subitems []models.Subitem
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&subitems)
	if db.Error != nil {
		return subitems, db.Error
	}
	return subitems, nil
}

//ALL ITEMS

func GetAllSubitems() ([]models.Subitem, error) {
	var subitems []models.Subitem
	db := database.DB.Preload(clause.Associations).Find(&subitems)
	if db.Error != nil {
		return subitems, db.Error
	}
	return subitems, nil
}
