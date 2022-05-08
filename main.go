package main

import (
	"github.com/kamil5b/backend-template/driver"
	"github.com/kamil5b/backend-template/utilities"
)

func main() {

	user := utilities.GoDotEnvVariable("DATABASE_USER")
	password := utilities.GoDotEnvVariable("DATABASE_PASSWORD")
	url := utilities.GoDotEnvVariable("DATABASE_URL")
	protocol := utilities.GoDotEnvVariable("DSN_PROTOCOL")
	db := utilities.GoDotEnvVariable("DATABASE_NAME")
	server := utilities.GoDotEnvVariable("SERVER_URL")
	//CREATE MORE THAN 1 SERVER GAIS!!!
	driver.SetupTemplate(server, url, user, password, protocol, db)
	//SetupTemplate("127.0.0.1:8080")
}
