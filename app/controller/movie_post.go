package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostMovie function
// @Summary      create a movie
// @Description  create a movie data
// @Tags         movies
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  model.Movie
// @Failure      400
// @Router       /movies [post]
func PostMovie(c *fiber.Ctx) error {
	db := services.DB
	var movieAPI model.MovieAPI
	if err := c.BodyParser(&movieAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	movie := new(model.Movie)
	movie.MovieAPI = movieAPI
	db.Create(&movie)

	return c.JSON(movie)
}
