package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetMovieID function
func GetMovieID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB
	var movie model.Movie
	result := db.Where("id = ?", id).First(&movie)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "movie not found",
		})
	}

	return c.JSON(movie)
}
