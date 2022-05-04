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
	app.Post("/api/user", controllers.PostUser)

	//----FORM----
	//CREATE
	app.Post("/api/form", controllers.PostForm)
	//READ
	app.Post("/api/form/id", controllers.GetFormByID)
	app.Post("/api/form/student", controllers.GetFormByStudent)
	app.Post("/api/form/teacher", controllers.GetFormByTeacher)
	app.Get("/api/form", controllers.GetAllForms)

	//----ITEM----
	//CREATE
	app.Post("/api/item", controllers.CreateItem)
	//READ
	app.Get("/api/items", controllers.GetItems)
	app.Post("/api/item/id", controllers.GetItemByID)
	app.Post("/api/item/name", controllers.GetItemsByName)
	//DELETE
	app.Delete("/api/item/id", controllers.DeleteItem)
	//UPDATE
	app.Put("/api/item/id", controllers.UpdateItem)

	//----SUBITEM----
	//CREATE
	app.Post("/api/subitem", controllers.CreateSubitem)
	//READ
	app.Get("/api/subitems", controllers.GetSubitems)
	app.Post("/api/subitem/id", controllers.GetSubitemByID)
	app.Post("/api/subitem/name", controllers.GetSubitemsByName)
	//DELETE
	app.Delete("/api/subitem/id", controllers.DeleteSubitem)
	//UPDATE
	app.Put("/api/subitem/id", controllers.UpdateSubitem)

}
