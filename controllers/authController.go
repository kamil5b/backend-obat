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
		Email string
		nik : string
		password : hash
		role : string
	*/

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	IP := c.IP()
	log := "history.log"
	msg := data["nik"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	dataint := utilities.MapStringToInt(data)
	user, err := repositories.CreateUser(data, dataint, IP)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	user, err = repositories.GetUser("ID = ?", user.ID)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	jsonClient, err := json.Marshal(user)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(jsonClient),
		ExpiresAt: time.Now().Add(time.Hour * 6).Unix(), //6 hours
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	msg = data["nik"] + " berhasil register unverified "

	utilities.WriteLog(log, IP, msg)

	// SEND EMAIL.
	/*
		url_verif := utilities.GoDotEnvVariable("SERVER_URL") + "/api/verify/" + token
		utilities.SendEmail(data["email"], "VERIFY", url_verif)
	*/
	c.Status(400)
	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func LoginUser(c *fiber.Ctx) error { //POST
	var data map[string]string
	//var nik string
	//IP bisa jadi server frontend atau backend atau client. Target : TRACE IP CLIENT
	/*

		nik : string
		password : string

	*/
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	IP := c.IP()
	log := "history.log"
	msg := data["nik"] + " mencoba untuk login"
	utilities.WriteLog(log, IP, msg)
	password := utilities.HashKamil(data["password"])
	user, err := repositories.GetUser("nik = ? and password = ? and verified", data["nik"], password)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}
	jsonClient, err := json.Marshal(user)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(jsonClient),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //6 hour
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		})
	}

	msg = data["nik"] + " berhasil login sebagai " + user.Role
	utilities.WriteLog(log, IP, msg)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
		"sessid":  token,
	})
}

func IsAuthorized(cookie, SecretKey string) bool {
	_, err := repositories.DecodeJWT(cookie, SecretKey)
	return err == nil
}

/*

func Login(data map[string]string, dataint map[string]int, IP string, SecretKey string) (string, int) {
	log := "history.log"
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := data["nik"] + " mencoba untuk login"
	utilities.WriteLog(log, IP, msg)
	password := utilities.HashKamil(data["password"])
	exist := IsUserExist("nik = ? and password = ?", data["nik"], password)
	if !exist {
		msg = data["nik"] + " belum mendaftar"
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": msg,
			"sessid":  nil,
		}, 400
	}
	user, err := GetAUser("nik = ? and password = ?", data["nik"], password)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		}, 400
	}
	jsonClient, err := json.Marshal(user)
	if err != nil {
		msg = data["nik"] + " " + err.Error()
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
		msg = data["nik"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return fiber.Map{
			"message": err.Error(),
			"sessid":  nil,
		}, 400
	}

	msg = data["nik"] + " berhasil login sebagai " + user.Role
	utilities.WriteLog(log, IP, msg)
	return fiber.Map{
		"message": "success",
		"sessid":  token,
	}, 200
}
*/
