package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

func DecodeJWT(cookie, SecretKey string) (models.User, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return models.User{}, err
	}
	var tmp models.User
	claims := token.Claims.(*jwt.StandardClaims)
	err = json.Unmarshal([]byte(claims.Issuer), &tmp)
	if err != nil {
		return models.User{}, err
	}
	return tmp, nil
}

//CREATE
func CreateUser(data map[string]string, dataint map[string]int, IP string) (models.User, error) {
	password := utilities.HashKamil(data["password"])
	log := utilities.GoDotEnvVariable("LOG")
	fmt.Println(data)
	exist := IsUserExist("username = ? or email = ?", data["username"], data["email"])
	if !exist {
		/*
			belum register. Register masukin ke table user hash
		*/
		user := models.User{
			Username: data["username"],
			Name:     data["name"],
			Password: password,
			Role:     data["role"],
			Email:    data["email"],
			Verified: false,
		}
		database.DB.Create(&user)
		msg := data["username"] + " berhasil mendaftar"
		utilities.WriteLog(log, IP, msg)
		fmt.Println(user)
		return user, nil
	}
	msg := data["username"] + " telah mendaftar sebelumnya"
	utilities.WriteLog(log, IP, msg)
	return models.User{}, errors.New(msg)
}

func VerifyUser(user models.User, IP string) error {
	log := utilities.GoDotEnvVariable("LOG")
	msg := strconv.Itoa(int(user.ID)) + " verify"
	utilities.WriteLog(log, IP, msg)
	fmt.Println(user)
	db := database.DB.Model(&user).Where("ID = ?", user.ID).Updates(models.User{
		Verified: true,
	})
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(int(user.ID)) + " verified successfully"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//DELETE
func DeleteUser(IP string, ID int) error { //DELETE

	log := utilities.GoDotEnvVariable("LOG")
	msg := strconv.Itoa(ID) + " menghapus"
	utilities.WriteLog(log, IP, msg)
	user, err := GetModelUser("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	db := database.DB.Delete(&user)
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil terhapus"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//UPDATE
func UpdateUser(data map[string]string, dataint map[string]int, IP string, ID int) error { //DELETE

	log := utilities.GoDotEnvVariable("LOG")
	msg := strconv.Itoa(ID) + " update"
	utilities.WriteLog(log, IP, msg)
	user, err := GetModelUser("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		return err
	}
	password := utilities.HashKamil(data["password"])
	db := database.DB.Model(&user).Updates(models.User{
		Username: data["username"],
		Name:     data["name"],
		Password: password,
		Role:     data["role"],
		Email:    data["email"],
	})
	if db.Error != nil {
		utilities.WriteLog(log, IP, db.Error.Error())
		return db.Error
	}

	msg = strconv.Itoa(ID) + " berhasil update"
	utilities.WriteLog(log, IP, msg)
	return nil
}

//READ

func IsUserExist(query string, val ...interface{}) bool {
	var user models.User
	db := database.DB.Where(query, val...).Last(&user)
	return db.Error == nil
}

func GetModelUser(query string, val ...interface{}) (models.User, error) {
	var user models.User
	db := database.DB.Where(query, val...).Last(&user)
	return user, db.Error
}

func GetModelUsers(query string, val ...interface{}) ([]models.User, error) {
	var users []models.User
	db := database.DB.Where(query, val...).Find(&users)
	return users, db.Error
}

func GetAllUser() ([]models.User, error) {
	var users []models.User
	db := database.DB.Find(&users)
	return users, db.Error
}
