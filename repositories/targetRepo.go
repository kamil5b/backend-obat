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

func CreateTargetSales(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := strconv.Itoa(dataint["int1"]) + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	user, err := GetUser("ID = ?", dataint["id_other"])
	if err != nil {
		msg = data["name"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return err

	} else if user.ID == 0 {
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
	database.DB.Create(&models.TargetSales{
		UserID:        user.ID,
		User:          user,
		NominalTarget: uint(dataint["nominal_target"]),
		TanggalTarget: tanggal,
	})
	msg = data["name"] + " berhasil dibuat"
	utilities.WriteLog(log, IP, msg)
	return nil

}

//=====DELETE=====
func DeleteTargetSales(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	target, err := GetTargetSales("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&target)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateTargetSales(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	target, err := GetTargetSales("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	user, err := GetUser("ID = ?", dataint["id_other"])
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	tanggal, err := utilities.ParsingDate(data["tanggal"])

	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}

	db := database.DB.Model(&target).Updates(models.TargetSales{
		UserID:        user.ID,
		User:          user,
		NominalTarget: uint(dataint["nominal_target"]),
		TanggalTarget: tanggal,
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

func GetTargetSales(query string, val ...interface{}) (models.TargetSales, error) {
	var target models.TargetSales
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&target)
	if db.Error != nil {
		return target, db.Error
	}
	return target, nil
}

//ARRAY

func GetTargetSaless(query string, val ...interface{}) ([]models.TargetSales, error) {
	var targets []models.TargetSales
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&targets)
	if db.Error != nil {
		return targets, db.Error
	}
	return targets, nil
}

//ALL ITEMS

func GetAllTargetSaless() ([]models.TargetSales, error) {
	var targets []models.TargetSales
	db := database.DB.Preload(clause.Associations).Find(&targets)
	if db.Error != nil {
		return targets, db.Error
	}
	return targets, nil
}
