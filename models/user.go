package models

import (
	"errors"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/utilities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

func IsUserExist(query string, val ...interface{}) bool {
	var user User
	database.DB.Where(query, val).First(&user)
	return user.ID == 0
}

func GetModelUser(query string, val ...interface{}) (User, error) {
	var user User
	db := database.DB.Where(query, val).First(&user)
	return user, db.Error
}

func GetAllUser(query string, val ...interface{}) ([]User, error) {
	var users []User
	db := database.DB.Where(query, val).Find(&users)
	return users, db.Error
}

func CreateUser(data map[string]string, dataint map[string]int, IP string) error {
	password := utilities.HashKamil(data["password"])
	log := utilities.GoDotEnvVariable("LOG")
	exist := IsUserExist("username = ?", data["username"])
	if !exist {
		/*
			belum register. Register masukin ke table user hash
		*/
		database.DB.Create(&User{
			Username: data["username"],
			Name:     data["name"],
			Password: password,
			Role:     data["role"],
		})
		msg := data["username"] + " berhasil mendaftar"
		utilities.WriteLog(log, IP, msg)
		return errors.New(msg)
	}
	msg := data["username"] + " telah mendaftar sebelumnya"
	utilities.WriteLog(log, IP, msg)
	return nil
}
