package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteMovie function
func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB
	var movie model.Movie
	result := db.Where("id = ?", id).First(&movie)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "movie not found",
		})
	}

	db.Where("id = ?", id).Delete(&movie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
