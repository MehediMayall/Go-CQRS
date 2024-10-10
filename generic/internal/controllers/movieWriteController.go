package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/handlers/commands"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type MovieWriteController struct {
	repo     repositories.IWriteRepository[entities.Movie]
	readRepo repositories.IReadRepository[entities.Movie]
}

func NewMovieController(
	repo repositories.IWriteRepository[entities.Movie],
	readRepo repositories.IReadRepository[entities.Movie]) MovieWriteController {

	return MovieWriteController{repo, readRepo}
}

// Save a Movie
func (c *MovieWriteController) Add(ctx *fiber.Ctx) error {
	movie := entities.Movie{}

	if err := ctx.BodyParser(&movie); err != nil {
		return BadRequest(ctx, err)
	}

	command := commands.NewAddMovieCommand(c.repo)
	movieId, err := command.Handle(&movie)
	if err != nil {
		return ServerError(ctx, err)
	}

	return Ok(ctx, movieId)
}

// Update a Movie
func (c *MovieWriteController) Update(ctx *fiber.Ctx) error {
	movie := entities.Movie{}

	if err := ctx.BodyParser(&movie); err != nil {
		return BadRequest(ctx, err)
	}

	command := commands.NewUpdateMovieCommand(c.repo, c.readRepo)
	err := command.Handle(&movie)
	if err != nil {
		return ServerError(ctx, err)
	}

	return Ok(ctx, "")
}

// Delete a Movie
func (c *MovieWriteController) Delete(ctx *fiber.Ctx) error {
	movieId := ctx.Params("id")

	if movieId == "" {
		return NotFound(ctx, errors.New("please provide a valid movie id"))
	}

	command := commands.NewDeleteMovieCommand(c.repo)
	err := command.Handle(movieId)
	if err != nil {
		return ServerError(ctx, err)
	}

	return Ok(ctx, "")
}
