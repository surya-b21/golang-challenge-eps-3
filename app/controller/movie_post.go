package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostMovie function
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
