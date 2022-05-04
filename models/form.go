package models

import (
	"errors"

	"github.com/kamil5b/backend-template/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Form struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	TeacherID int    `json:"id_teacher"`
	Teacher   User   `json:"teacher"`
	StudentID int    `json:"id_student" gorm:"unique"`
	Student   User   `json:"student"`
	Input     string `json:"input"`
	UrlVoice  string `json:"url_voice" gorm:"unique"`
	UrlImage  string `json:"url_image" gorm:"unique"`
}

func CreateForm(data map[string]string, dataint map[string]int, IP string) error { //POST
	teacher, err := GetModelUser("ID = ?", dataint["id_teacher"])
	if err != nil {
		return err
	} else if teacher.ID == 0 {
		return errors.New("teacher doesn't exist")
	}
	//GET STUDENT
	student, err := GetModelUser("ID = ?", dataint["id_student"])
	if err != nil {
		return err
	} else if student.ID == 0 {
		return errors.New("student doesn't exist")
	}
	//CREATE FORM
	database.DB.Create(&Form{
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
func GetForm(id int) (Form, error) {
	var form Form
	db := database.DB.Where("id = ?", id).Preload(clause.Associations).First(&form)

	return form, db.Error
}

func GetFormStudent(id int) ([]Form, error) {
	var forms []Form
	db := database.DB.Where("id_student = ?", id).Preload(clause.Associations).Find(&forms)
	return forms, db.Error
}

func GetFormTeacher(id int) ([]Form, error) {
	var forms []Form
	db := database.DB.Where("id_teacher = ?", id).Preload(clause.Associations).Find(&forms)
	return forms, db.Error
}

func GetAllForms() ([]Form, error) {
	var forms []Form
	db := database.DB.Preload(clause.Associations).Find(&forms)
	return forms, db.Error
}
