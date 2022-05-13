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

func CreateGiro(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := strconv.Itoa(dataint["int1"]) + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	bank, err := GetBank("NomorBank = ?", dataint["nomor_bank"])
	if err != nil {
		msg = data["name"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return err

	} else if bank.ID == 0 {
		//CREATE FORM
		msg = strconv.Itoa(dataint["id_other"]) + " tidak dbankukan"
		utilities.WriteLog(log, IP, msg)
		return errors.New(msg)
	}

	/*
		belum register. Register masukin ke table user hash
	*/

	tanggal, err := utilities.ParsingDate(data["tanggal"])

	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	database.DB.Create(&models.Giro{
		Tanggal:   tanggal,
		NomorGiro: data["nomor_giro"],
		BankID:    bank.ID,
		Bank:      bank,
		Nominal:   uint(dataint["nominal"]),
	})
	msg = data["name"] + " berhasil dibuat"
	utilities.WriteLog(log, IP, msg)
	return nil

}

//=====DELETE=====
func DeleteGiro(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	giro, err := GetGiro("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&giro)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateGiro(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	giro, err := GetGiro("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	bank, err := GetBank("ID = ?", dataint["id_other"])
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	tanggal, err := utilities.ParsingDate(data["tanggal"])

	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&giro).Updates(models.Giro{
		Tanggal:   tanggal,
		NomorGiro: data["nomor_giro"],
		BankID:    bank.ID,
		Bank:      bank,
		Nominal:   uint(dataint["nominal"]),
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

func GetGiro(query string, val ...interface{}) (models.Giro, error) {
	var giro models.Giro
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&giro)
	if db.Error != nil {
		return giro, db.Error
	}
	return giro, nil
}

//ARRAY

func GetGiros(query string, val ...interface{}) ([]models.Giro, error) {
	var giros []models.Giro
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&giros)
	if db.Error != nil {
		return giros, db.Error
	}
	return giros, nil
}

//ALL ITEMS

func GetAllGiros() ([]models.Giro, error) {
	var giros []models.Giro
	db := database.DB.Preload(clause.Associations).Find(&giros)
	if db.Error != nil {
		return giros, db.Error
	}
	return giros, nil
}
