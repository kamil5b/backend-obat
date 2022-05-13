package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateBarang(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.Barang{
		Kode:          data["kode"],
		Name:          data["name"],
		TipeBig:       data["tipe_big"],
		TipeMedium:    data["tipe_medium"],
		TipeSmall:     data["tipe_small"],
		BigToMedium:   uint(dataint["big_to_small"]),
		MediumToSmall: uint(dataint["medium_to_small"]),
		HargaBig:      data["harga_big"],
		HargaMedium:   data["harga_medium"],
		HargaSmall:    data["harga_small"],
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
func DeleteBarang(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	barang, err := GetBarang("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&barang)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateBarang(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	barang, err := GetBarang("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&barang).Updates(models.Barang{
		Kode:          data["kode"],
		Name:          data["name"],
		TipeBig:       data["tipe_big"],
		TipeMedium:    data["tipe_medium"],
		TipeSmall:     data["tipe_small"],
		BigToMedium:   uint(dataint["big_to_small"]),
		MediumToSmall: uint(dataint["medium_to_small"]),
		HargaBig:      data["harga_big"],
		HargaMedium:   data["harga_medium"],
		HargaSmall:    data["harga_small"],
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

func IsBarangExist(query string, val ...interface{}) bool {
	var barang models.Barang
	database.DB.Where(query, val...).Last(&barang)
	return barang.ID != 0
}

//GET AN ITEM

func GetBarang(query string, val ...interface{}) (models.Barang, error) {
	var barang models.Barang
	db := database.DB.Where(query, val...).Last(&barang)
	if db.Error != nil {
		return barang, db.Error
	}
	return barang, nil
}

//ARRAY

func GetBarangs(query string, val ...interface{}) ([]models.Barang, error) {
	var barangs []models.Barang
	db := database.DB.Where(query, val...).Find(&barangs)
	if db.Error != nil {
		return nil, db.Error
	}
	return barangs, nil
}

//ALL ITEMS

func GetAllBarangs() ([]models.Barang, error) {
	var barangs []models.Barang
	db := database.DB.Find(&barangs)
	if db.Error != nil {
		return nil, db.Error
	}
	return barangs, nil
}
