package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
	"github.com/kamil5b/backend-template/utilities"
)

func GetSubitems(c *fiber.Ctx) error { //GET
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	msg := "Get all item"
	utilities.WriteLog(log, IP, msg)
	items, err := repositories.GetAllSubitems()
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

func CreateSubitem(c *fiber.Ctx) error { //POST
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
	err := repositories.CreateSubitem(data, dataint, IP)
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

func DeleteSubitem(c *fiber.Ctx) error { //DELETE
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	err := repositories.DeleteSubitem(data, dataint, IP)
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

func UpdateSubitem(c *fiber.Ctx) error { //UPDATE
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	err := repositories.UpdateSubitem(data, dataint, IP)
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

func GetSubitemByID(c *fiber.Ctx) error { //POST
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
	item, err := repositories.GetASubitem("ID = ?", dataint["id"])
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

func GetSubitemsByName(c *fiber.Ctx) error { //POST
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
	items, err := repositories.GetASubitem("name = ?", data["name"])
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