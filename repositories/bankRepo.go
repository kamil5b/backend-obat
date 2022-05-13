package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateBank(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.Bank{
		Nomor: uint(dataint["nomor_bank"]),
		Nama:  data["nama_bank"],
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
func DeleteBank(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	bank, err := GetBank("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&bank)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateBank(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	bank, err := GetBank("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&bank).Updates(models.Bank{
		Nomor: uint(dataint["nomor_bank"]),
		Nama:  data["nama_bank"],
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

func IsBankExist(query string, val ...interface{}) bool {
	var bank models.Bank
	database.DB.Where(query, val...).Last(&bank)
	return bank.ID != 0
}

//GET AN ITEM

func GetBank(query string, val ...interface{}) (models.Bank, error) {
	var bank models.Bank
	db := database.DB.Where(query, val...).Last(&bank)
	if db.Error != nil {
		return bank, db.Error
	}
	return bank, nil
}

//ARRAY

func GetBanks(query string, val ...interface{}) ([]models.Bank, error) {
	var banks []models.Bank
	db := database.DB.Where(query, val...).Find(&banks)
	if db.Error != nil {
		return nil, db.Error
	}
	return banks, nil
}

//ALL ITEMS

func GetAllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	db := database.DB.Find(&banks)
	if db.Error != nil {
		return nil, db.Error
	}
	return banks, nil
}
