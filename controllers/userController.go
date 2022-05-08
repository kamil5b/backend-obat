package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
	"github.com/kamil5b/backend-template/utilities"
)

//CREATE
func AuthUser(c *fiber.Ctx) error { //POST
	auth := c.Params("auth")
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	user, err := repositories.DecodeJWT(auth, SecretKey)
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

func RegisterAuth(c *fiber.Ctx) error {
	auth := c.Params("auth")
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	user, err := repositories.DecodeJWT(auth, SecretKey)
	if err != nil {
		utilities.WriteLog(log, IP, "Gagal verified")
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = repositories.VerifyUser(user, IP)
	if err != nil {
		utilities.WriteLog(log, IP, "Gagal verify")
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

//DELETE

func DeleteUser(c *fiber.Ctx) error { //DELETE
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"form":    nil,
		})
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
	err = repositories.DeleteUser(IP, ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	utilities.WriteLog(log, IP, "Berhasil")
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

//UPDATE
func UpdateUser(c *fiber.Ctx) error { //UPDATE
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	dataint := utilities.MapStringToInt(data)

	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
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
	err = repositories.UpdateUser(data, dataint, IP, ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	utilities.WriteLog(log, IP, "Berhasil")
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

//READ
func GetUserByID(c *fiber.Ctx) error { //POST
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	ID, err := c.ParamsInt("id")
	if err != nil {
		utilities.WriteLog(log, IP, "fail")
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := repositories.GetUser("ID = ?", ID)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
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

func GetUserByName(c *fiber.Ctx) error { //POST
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	name, err := c.ParamsInt("name")
	if err != nil {
		utilities.WriteLog(log, IP, "fail")
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	users, err := repositories.GetUsers("name like %?%", name)
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	utilities.WriteLog(log, IP, "Berhasil")
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
		"users":   users,
	})
}

func GetUsers(c *fiber.Ctx) error { //POST
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	users, err := repositories.GetAllUser()
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	utilities.WriteLog(log, IP, "Berhasil")
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "success",
		"users":   users,
	})
}
