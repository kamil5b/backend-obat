package repositories

import (
	"strconv"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

//=====POST======

func CreateCustomer(data map[string]string, dataint map[string]int, IP string) error { //POST

	log := "history.log"
	msg := data["name"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	db := database.DB.Create(&models.Customer{
		Name:         data["name"],
		Alamat:       data["alamat"],
		NomorHP:      data["nomor_hp"],
		LimitPiutang: uint(dataint["limit_piutang"]),
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
func DeleteCustomer(IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	customer, err := GetCustomer("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&customer)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//=====UPDATE=====
func UpdateCustomer(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := "history.log"
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	customer, err := GetCustomer("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Model(&customer).Updates(models.Customer{
		Name:         data["name"],
		Alamat:       data["alamat"],
		NomorHP:      data["nomor_hp"],
		LimitPiutang: uint(dataint["limit_piutang"]),
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

func IsCustomerExist(query string, val ...interface{}) bool {
	var customer models.Customer
	database.DB.Where(query, val...).Last(&customer)
	return customer.ID != 0
}

//GET AN ITEM

func GetCustomer(query string, val ...interface{}) (models.Customer, error) {
	var customer models.Customer
	db := database.DB.Where(query, val...).Last(&customer)
	if db.Error != nil {
		return customer, db.Error
	}
	return customer, nil
}

//ARRAY

func GetCustomers(query string, val ...interface{}) ([]models.Customer, error) {
	var customers []models.Customer
	db := database.DB.Where(query, val...).Find(&customers)
	if db.Error != nil {
		return nil, db.Error
	}
	return customers, nil
}

//ALL ITEMS

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	db := database.DB.Find(&customers)
	if db.Error != nil {
		return nil, db.Error
	}
	return customers, nil
}
