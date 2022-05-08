package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
	"github.com/kamil5b/backend-template/utilities"
)

func GetItems(c *fiber.Ctx) error { //GET
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {

		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := "Get all item"
	utilities.WriteLog(log, IP, msg)
	items, err := repositories.GetAllItems()
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
		"items":   items,
	})
}

func CreateItem(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		value : hash
	*/
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	err := repositories.CreateItem(data, dataint, IP)
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

func DeleteItem(c *fiber.Ctx) error { //DELETE
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
	err = repositories.DeleteItem(IP, ID)
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

func UpdateItem(c *fiber.Ctx) error { //UPDATE
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
	err = repositories.UpdateItem(data, dataint, IP, ID)
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

func GetItemByID(c *fiber.Ctx) error { //POST
	/*
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return err
		}*/
	ID, err := c.ParamsInt("id")
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	//IP := c.IP()
	item, err := repositories.GetItem("ID = ?", ID)
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
		"item":    item,
	})
}

func GetItemsByName(c *fiber.Ctx) error { //POST
	name, err := c.ParamsInt("name")
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	items, err := repositories.GetItems("name like %?%", name)
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
		"items":   items,
	})
}
