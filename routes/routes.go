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
	app.Post("/api/:auth", controllers.PostUser)

	//----FORM----
	//CREATE
	app.Post("/api/form/:auth", controllers.PostForm)
	//READ
	app.Post("/api/form/id/:auth", controllers.GetFormByID)
	app.Post("/api/form/student/:auth", controllers.GetFormByStudent)
	app.Post("/api/form/teacher/:auth", controllers.GetFormByTeacher)
	app.Get("/api/form/:auth", controllers.GetAllForms)

	//----ITEM----
	//CREATE
	app.Post("/api/item/:auth", controllers.CreateItem)
	//READ
	app.Get("/api/items/:auth", controllers.GetItems)
	app.Post("/api/item/id/:auth", controllers.GetItemByID)
	app.Post("/api/item/name/:auth", controllers.GetItemsByName)
	//DELETE
	app.Delete("/api/item/id/:auth", controllers.DeleteItem)
	//UPDATE
	app.Put("/api/item/id/:auth", controllers.UpdateItem)

	//----SUBITEM----
	//CREATE
	app.Post("/api/subitem/:auth", controllers.CreateSubitem)
	//READ
	app.Get("/api/subitems/:auth", controllers.GetSubitems)
	app.Post("/api/subitem/id/:auth", controllers.GetSubitemByID)
	app.Post("/api/subitem/name/:auth", controllers.GetSubitemsByName)
	//DELETE
	app.Delete("/api/subitem/id/:auth", controllers.DeleteSubitem)
	//UPDATE
	app.Put("/api/subitem/id/:auth", controllers.UpdateSubitem)

}
