package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mehedimayall/go-cqrs/internal/handlers/queries"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type MovieReadController struct {
	repo repositories.IMovieReadRepository
}

func NewMovieReadController(repo repositories.IMovieReadRepository) MovieReadController {
	return MovieReadController{
		repo: repo,
	}
}

func (c *MovieReadController) GetAll(ctx *fiber.Ctx) error {
	query := queries.NewGetMoviesQuery(c.repo)
	result, err := query.Handle()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(result)
}
