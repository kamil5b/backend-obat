package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/routes"
	"github.com/kamil5b/backend-template/utilities"
)

func SetupTemplate(serverurl string) {
	database.Connect()
	app := fiber.New()
	/*
		origin := utilities.GoDotEnvVariable("VIEW_URL") //ganti view url ini di .env
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		}))
	*/
	routes.Setup(app)

	err := app.Listen(serverurl)
	if err != nil {
		fmt.Println(err)
		fmt.Scan()
	}
}

func main() {
	//CREATE MORE THAN 1 SERVER GAIS!!!
	go SetupTemplate(utilities.GoDotEnvVariable("SERVER_URL"))
	SetupTemplate("127.0.0.1:8080")
}
