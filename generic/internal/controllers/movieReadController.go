package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/handlers/queries"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type MovieReadController struct {
	repo repositories.IReadRepository[entities.Movie]
}

func NewMovieReadController(repo repositories.IReadRepository[entities.Movie]) MovieReadController {
	return MovieReadController{repo}
}

// GET ALL MOVIES
func (c *MovieReadController) GetAll(ctx *fiber.Ctx) error {
	query := queries.NewGetMoviesQuery(c.repo)
	result, err := query.Handle()

	if err != nil {
		return ServerError(ctx, err)
	}

	return Ok(ctx, result)
}

// GET A MOVIE
func (c *MovieReadController) GetById(ctx *fiber.Ctx) error {
	movieId := ctx.Params("id")
	if movieId == "" {
		return NotFound(ctx, errors.New("please provide a valid movie id"))
	}

	query := queries.NewGetMovieByIdQuery(c.repo)
	result, err := query.Handle(movieId)

	if err != nil {
		return ServerError(ctx, err)
	}

	return Ok(ctx, result)
}
