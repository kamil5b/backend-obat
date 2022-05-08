package repositories

import (
	"errors"

	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"gorm.io/gorm/clause"
)

func CreateForm(data map[string]string, dataint map[string]int, IP string) error { //POST
	teacher, err := GetUser("ID = ?", dataint["id_teacher"])
	if err != nil {
		return err
	} else if teacher.ID == 0 {
		return errors.New("teacher doesn't exist")
	}
	//GET STUDENT
	student, err := GetUser("ID = ?", dataint["id_student"])
	if err != nil {
		return err
	} else if student.ID == 0 {
		return errors.New("student doesn't exist")
	}
	//CREATE FORM
	database.DB.Create(&models.Form{
		TeacherID: dataint["id_teacher"],
		Teacher:   teacher,
		StudentID: dataint["id_student"],
		Student:   student,
		Input:     data["text"],
		UrlVoice:  data["url_voice"],
		UrlImage:  data["url_image"],
	})
	return nil
}

//GET AN ITEM

func GetForm(query string, val ...interface{}) (models.Form, error) {
	var form models.Form
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&form)
	if db.Error != nil {
		return form, db.Error
	}
	return form, nil
}

//ARRAY

func GetForms(query string, val ...interface{}) ([]models.Form, error) {
	var forms []models.Form
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&forms)
	if db.Error != nil {
		return forms, db.Error
	}
	return forms, nil
}

func GetAllForms() ([]models.Form, error) {
	var forms []models.Form
	db := database.DB.Preload(clause.Associations).Find(&forms)
	return forms, db.Error
}
