package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(url, user, password, protocol, database string, models ...interface{}) {

	dsn := user + ":" + password + "@" + protocol + "(" + url + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(models...)
}
