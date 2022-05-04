package repositories

import (
	"encoding/json"
	"errors"

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

func IsUserExist(query string, val ...interface{}) bool {
	var user models.User
	database.DB.Where(query, val).Last(&user)
	return user.ID == 0
}

func GetModelUser(query string, val ...interface{}) (models.User, error) {
	var user models.User
	db := database.DB.Where(query, val).Last(&user)
	return user, db.Error
}

func GetAllUser(query string, val ...interface{}) ([]models.User, error) {
	var users []models.User
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
		database.DB.Create(&models.User{
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
