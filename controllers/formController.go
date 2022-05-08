package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/repositories"
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
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	recording, err := c.FormFile("wavfile")
	if err != nil {
		return err
	}

	image, err := c.FormFile("image")
	if err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	IP := c.IP()
	log := "history.log"
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	msg := data["username"] + " mendaftar"
	utilities.WriteLog(log, IP, msg)

	data["url_voice"] = fmt.Sprintf("../storage/record/%s", recording.Filename)
	data["url_image"] = fmt.Sprintf("../storage/images/%s", image.Filename)
	err = repositories.CreateForm(data, dataint, IP)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		msg = data["id_student"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//SAVE VOICE
	err = c.SaveFile(recording, data["url_voice"])
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		msg = data["id_student"] + " " + err.Error()
		utilities.WriteLog(log, IP, msg)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//SAVE IMAGE
	err = c.SaveFile(image, data["url_image"])
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
	forms, err := repositories.GetAllForms()
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
	ID, err := c.ParamsInt("id")
	IP := c.IP()
	log := "history.log"
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	form, err := repositories.GetForm("ID = ?", ID)
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

	ID, err := c.ParamsInt("id")
	IP := c.IP()
	log := "history.log"
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if err != nil {
		utilities.WriteLog(log, IP, err.Error())
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	forms, err := repositories.GetForms("id_student = ?", ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"form":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"forms":   forms,
	})

}

func GetFormByTeacher(c *fiber.Ctx) error { //POST

	ID, err := c.ParamsInt("id")
	IP := c.IP()
	log := "history.log"
	if !IsAuthorized(c.Params("auth"), SecretKey) {
		utilities.WriteLog(log, IP, "unauthorized")
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"forms":   nil,
		})
	}
	forms, err := repositories.GetForms("id_teacher = ?", ID)
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
