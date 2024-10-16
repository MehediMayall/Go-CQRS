package app

import (
	"github.com/mehedimayall/go-cqrs/internal/controllers"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

func (app *App) CreateWriteMovieController() controllers.MovieWriteController {
	movieRepo := repositories.NewMovieWriteRepository(app.db)
	return controllers.NewMovieController(movieRepo)
}

func (app *App) CreateReadMovieController() controllers.MovieReadController {
	repo := repositories.NewMovieReadRepository(app.db)
	return controllers.NewMovieReadController(repo)
}
