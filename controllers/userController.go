package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
	"github.com/kamil5b/backend-template/utilities"
)

func PostUser(c *fiber.Ctx) error { //POST
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	user, err := repositories.DecodeJWT(data["value"], SecretKey)
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
