package models

import (
	"gorm.io/gorm"
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
