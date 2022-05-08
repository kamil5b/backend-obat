package template

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/database"
	"github.com/kamil5b/backend-template/models"
	"github.com/kamil5b/backend-template/routes"
)

func SetupTemplate(server_url, db_url, user, password, protocol, db string) {

	database.Connect(
		db_url, user, password, protocol, db,
		&models.Item{},
		&models.Subitem{},
		&models.User{},
		&models.Form{},
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

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}
