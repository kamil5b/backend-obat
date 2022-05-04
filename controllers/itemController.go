package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

func GetItems(c *fiber.Ctx) error { //GET
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	msg := "Get all item"
	utilities.WriteLog(log, IP, msg)
	items, err := models.GetAllItems()
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

func RegisterItem(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		value : hash
	*/
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	err := models.CreateItem(data, dataint, IP)
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
	var data map[string]string
	/*
		value : hash
	*/
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	//IP := c.IP()
	item, err := models.GetAnItem("ID = ?", dataint["id"])
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
		"items":   item,
	})
}

func GetItemsByName(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		value : hash
	*/
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//IP := c.IP()
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	items, err := models.GetAnItem("name = ?", data["name"])
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
