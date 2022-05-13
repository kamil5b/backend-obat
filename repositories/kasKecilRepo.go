package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateKasKecil(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.KasKecil{
		Nominal:     dataint["nominal"],
		Description: dataint["description"],
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
func DeleteKasKecil(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	kas, err := GetKasKecil("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&kas)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateKasKecil(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	kas, err := GetKasKecil("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&kas).Updates(models.KasKecil{
		Nominal:     dataint["nominal"],
		Description: dataint["description"],
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

func IsKasKecilExist(query string, val ...interface{}) bool {
	var kas models.KasKecil
	database.DB.Where(query, val...).Last(&kas)
	return kas.ID != 0
}

//GET AN ITEM

func GetKasKecil(query string, val ...interface{}) (models.KasKecil, error) {
	var kas models.KasKecil
	db := database.DB.Where(query, val...).Last(&kas)
	if db.Error != nil {
		return kas, db.Error
	}
	return kas, nil
}

//ARRAY

func GetKasKecils(query string, val ...interface{}) ([]models.KasKecil, error) {
	var kass []models.KasKecil
	db := database.DB.Where(query, val...).Find(&kass)
	if db.Error != nil {
		return nil, db.Error
	}
	return kass, nil
}

//ALL ITEMS

func GetAllKasKecils() ([]models.KasKecil, error) {
	var kass []models.KasKecil
	db := database.DB.Find(&kass)
	if db.Error != nil {
		return nil, db.Error
	}
	return kass, nil
}
