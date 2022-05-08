package database

import (
	"github.com/kamil5b/backend-template/utilities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(models ...interface{}) {
	user := utilities.GoDotEnvVariable("DATABASE_USER")
	password := utilities.GoDotEnvVariable("DATABASE_PASSWORD")
	url := utilities.GoDotEnvVariable("DATABASE_URL")
	protocol := utilities.GoDotEnvVariable("DSN_PROTOCOL")
	database := utilities.GoDotEnvVariable("DATABASE_NAME")

	dsn := user + ":" + password + "@" + protocol + "(" + url + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(models...)
}
