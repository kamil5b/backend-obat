package controllers

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
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

func PostUser(c *fiber.Ctx) error { //POST
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	user, err := DecodeJWT(data["value"], SecretKey)
	if err != nil {
		utilities.WriteLog(log, IP, "Gagal login")
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	utilities.WriteLog(log, IP, "Berhasil")
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}
