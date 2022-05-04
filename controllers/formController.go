package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/utilities"
)

func PostForm(c *fiber.Ctx) error { //POST

	/*
		id_teacher : int
		id_student : int
		url_voice : string
		text : string
		url_image : string


		 file, err := c.FormFile("document")

		// Save file to root directory:
		return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	*/
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	recording, err := c.FormFile("wavfile")
	if err != nil {
		return err
	}

	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := utilities.GoDotEnvVariable("LOG")
	msg := data["username"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)
	data["url_voice"] = fmt.Sprintf("../storage/record/%s", recording.Filename)
	err = models.CreateForm(data, dataint, IP)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		msg = data["id_student"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = c.SaveFile(recording, data["url_voice"])
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		msg = data["id_student"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	msg = data["id_student"] + " berhasil form"
	utilities.WriteLog(log, IP, msg)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func GetAllForms(c *fiber.Ctx) error { //GET
	forms, err := models.GetAllForms()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"forms":   nil,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"forms":   forms,
	})
}

func GetFormByID(c *fiber.Ctx) error { //POST
	/*
		id : int
	*/
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)

	form, err := models.GetForm(dataint["id"])
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"form":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"form":    form,
	})

}

func GetFormByStudent(c *fiber.Ctx) error { //POST
	/*
		id_student : int
	*/
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)

	form, err := models.GetFormStudent(dataint["id_student"])
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"form":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"form":    form,
	})

}

func GetFormByTeacher(c *fiber.Ctx) error { //POST
	/*
		id_teacher : int
	*/
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)

	forms, err := models.GetFormTeacher(dataint["id_teacher"])
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"forms":   nil,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"forms":   forms,
	})

}
