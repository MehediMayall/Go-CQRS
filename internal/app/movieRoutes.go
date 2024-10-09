package app

import "github.com/gofiber/fiber/v2"

func (app *App) RegisterMovieRoutes(api *fiber.App) {

	// Version
	v1 := api.Group("/movie")

	controller := app.CreateWriteMovieController()
	v1.Get("", controller.Add)
}
