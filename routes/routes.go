package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamil5b/backend-template/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})

	//----AUTH----
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.LoginUser)
	app.Get("/api/:auth", controllers.AuthUser)
	app.Get("/api/verify/:auth", controllers.RegisterAuth)

	//----USER----
	app.Get("/api/users/:auth", controllers.GetUsers)
	app.Get("/api/user/:id/:auth", controllers.GetUserByID)
	app.Get("/api/users/:name/:auth", controllers.GetUserByName)
	//DELETE
	app.Delete("/api/user/:id/:auth", controllers.DeleteUser)
	//UPDATE
	app.Put("/api/user/:id/:auth", controllers.UpdateUser)

	//----FORM----
	//CREATE
	app.Post("/api/form/:auth", controllers.PostForm)
	//READ
	app.Get("/api/form/:id/:auth", controllers.GetFormByID)
	app.Get("/api/form/student/:id/:auth", controllers.GetFormByStudent)
	app.Get("/api/form/teacher/:id/:auth", controllers.GetFormByTeacher)
	app.Get("/api/form/:auth", controllers.GetAllForms)

	//----ITEM----
	//CREATE
	app.Post("/api/item/:auth", controllers.CreateItem)
	//READ
	app.Get("/api/items/:auth", controllers.GetItems)
	app.Get("/api/item/:id/:auth", controllers.GetItemByID)
	app.Get("/api/items/:name/:auth", controllers.GetItemsByName)
	//DELETE
	app.Delete("/api/item/:id/:auth", controllers.DeleteItem)
	//UPDATE
	app.Put("/api/item/:id/:auth", controllers.UpdateItem)

	//----SUBITEM----
	//CREATE
	app.Post("/api/subitem/:auth", controllers.CreateSubitem)
	//READ
	app.Get("/api/subitems/:auth", controllers.GetSubitems)
	app.Get("/api/subitem/:id/:auth", controllers.GetSubitemByID)
	app.Get("/api/subitems/:name/:auth", controllers.GetSubitemsByName)
	//DELETE
	app.Delete("/api/subitem/:id/:auth", controllers.DeleteSubitem)
	//UPDATE
	app.Put("/api/subitem/:id/:auth", controllers.UpdateSubitem)

}
