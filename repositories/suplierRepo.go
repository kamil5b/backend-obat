package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateSuplier(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.Suplier{
		Name:    data["name"],
		Alamat:  data["alamat"],
		NomorHP: data["nomor_hp"],
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
func DeleteSuplier(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	suplier, err := GetSuplier("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&suplier)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateSuplier(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	suplier, err := GetSuplier("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&suplier).Updates(models.Suplier{
		Name:    data["name"],
		Alamat:  data["alamat"],
		NomorHP: data["nomor_hp"],
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

func IsSuplierExist(query string, val ...interface{}) bool {
	var suplier models.Suplier
	database.DB.Where(query, val...).Last(&suplier)
	return suplier.ID != 0
}

//GET AN ITEM

func GetSuplier(query string, val ...interface{}) (models.Suplier, error) {
	var suplier models.Suplier
	db := database.DB.Where(query, val...).Last(&suplier)
	if db.Error != nil {
		return suplier, db.Error
	}
	return suplier, nil
}

//ARRAY

func GetSupliers(query string, val ...interface{}) ([]models.Suplier, error) {
	var supliers []models.Suplier
	db := database.DB.Where(query, val...).Find(&supliers)
	if db.Error != nil {
		return nil, db.Error
	}
	return supliers, nil
}

//ALL ITEMS

func GetAllSupliers() ([]models.Suplier, error) {
	var supliers []models.Suplier
	db := database.DB.Find(&supliers)
	if db.Error != nil {
		return nil, db.Error
	}
	return supliers, nil
}
