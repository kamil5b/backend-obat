package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/routes"
	"github.com/kamil5b/backend-template/utilities"
)

func main() {
	database.Connect(
		&models.Item,
		&models.Subitem,
		&models.User,
	)
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
	serverurl := utilities.GoDotEnvVariable("SERVER_URL")

	err := app.Listen(serverurl)
	if err != nil {
		fmt.Println(err)
		fmt.Scan()
	}
}
