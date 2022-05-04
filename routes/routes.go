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
	//app.Post("/api/user", controllers.PostUser)

	//----FORM----
	app.Post("/api/form", controllers.PostForm)
	app.Post("/api/form/id", controllers.GetFormByID)
	app.Post("/api/form/student", controllers.GetFormByStudent)
	app.Post("/api/form/teacher", controllers.GetFormByTeacher)
	app.Get("/api/form", controllers.GetAllForms)
}
