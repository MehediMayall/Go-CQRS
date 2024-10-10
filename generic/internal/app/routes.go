package app

import "github.com/gofiber/fiber/v2"

func (app *App) RegisterMovieRoutes(api fiber.Router) {

	// Version
	v1 := api.Group("/")
	mg := v1.Group("movies")

	// Write
	controller := app.CreateWriteMovieController()
	mg.Post("", controller.Add)
	mg.Put("", controller.Update)
	mg.Delete("/:id", controller.Delete)

	// Read
	readController := app.CreateReadMovieController()
	mg.Get("", readController.GetAll)

	mg.Get("/:id", readController.GetById)
}
