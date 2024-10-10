package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehedimayall/go-cqrs/internal/database"
)

type App struct {
	db *database.InMemoryDB
}

func Run() {
	app := new(App)

	db := database.NewDB()

	app.db = db

	// Server
	server := fiber.New()
	api := server.Group("/api")

	// Register Routes
	app.RegisterMovieRoutes(api)

	// Start Server
	server.Listen(":3500")

}
