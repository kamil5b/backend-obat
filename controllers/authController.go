package controllers

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
	"github.com/kamil5b/backend-template/utilities"
)

const SecretKey = "314221943871221"

func RegisterUser(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		username : string
		password : hash
		role : string
	*/

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := data["username"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	dataint := utilities.MapStringToInt(data)
	err := repositories.CreateUser(data, dataint, IP)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func LoginUser(c *fiber.Ctx) error { //POST
	var data map[string]string
	//var username string
	//IP bisa jadi server frontend atau backend atau client. Target : TRACE IP CLIENT
	/*

		username : string
		password : string

	*/
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := data["username"] + " mencoba untuk login"
	utilities.WriteLog(log, IP, msg)
	password := utilities.HashKamil(data["password"])
	user, err := repositories.GetModelUser("username = ? and password = ?", data["username"], password)
	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}
	jsonClient, err := json.Marshal(user)
	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(jsonClient),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}

	msg = data["username"] + " berhasil login sebagai " + user.Role
	utilities.WriteLog(log, IP, msg)
	c.Status(400)
	return c.JSON(fiber.Map{
		"message": "success",
		"sessid":  token,
	})
}

func IsAuthorized(cookie, SecretKey string) bool {
	_, err := repositories.DecodeJWT(cookie, SecretKey)
	return err != nil
}

/*

func Login(data map[string]string, dataint map[string]int, IP string, SecretKey string) (string, int) {
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := data["username"] + " mencoba untuk login"
	utilities.WriteLog(log, IP, msg)
	password := utilities.HashKamil(data["password"])
	exist := IsUserExist("username = ? and password = ?", data["username"], password)
	if !exist {
		msg = data["username"] + " belum mendaftar"
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": msg,
			"sessid":  nil,
		}, 400
	}
	user, err := GetAUser("username = ? and password = ?", data["username"], password)
	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		}, 400
	}
	jsonClient, err := json.Marshal(user)
	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		}, 400
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(jsonClient),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		msg = data["username"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		}, 400
	}

	msg = data["username"] + " berhasil login sebagai " + user.Role
	utilities.WriteLog(log, IP, msg)
	return fiber.Map{
		"message": "success",
		"sessid":  token,
	}, 200
}
*/
