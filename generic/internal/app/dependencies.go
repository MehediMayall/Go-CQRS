// Manual Dependency Injection. This go file manages the dependencies of the application.
// Every functions in this file is used to instantiate the dependent objects and
// inject them into the parent object through constructor injection

package app

import (
	"github.com/mehedimayall/go-cqrs/internal/controllers"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

func (app *App) CreateWriteMovieController() controllers.MovieWriteController {
	movieRepo := repositories.NewMovieWriteRepository(app.db)
	movieReadRepo := repositories.NewMovieReadRepository(app.db)
	return controllers.NewMovieController(movieRepo, &movieReadRepo)
}

func (app *App) CreateReadMovieController() controllers.MovieReadController {
	repo := repositories.NewMovieReadRepository(app.db)
	return controllers.NewMovieReadController(&repo)
}
