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
	server.Group("/api")

	// Register Routes
	app.RegisterMovieRoutes(server)

	// Start Server
	server.Listen(":3500")

}
