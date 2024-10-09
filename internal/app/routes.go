package app

import "github.com/gofiber/fiber/v2"

func (app *App) RegisterMovieRoutes(server *fiber.App) {
	api := server.Group("/api")

	// Version
	v1 := api.Group("/")
	mg := v1.Group("movies")

	controller := app.CreateWriteMovieController()
	mg.Post("", controller.Add)

	readController := app.CreateReadMovieController()
	mg.Get("", readController.GetAll)
}
