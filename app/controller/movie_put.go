package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// PutMovie function
func PutMovie(c *fiber.Ctx) error {
	db := services.DB
	var movieAPI model.MovieAPI
	if err := c.BodyParser(&movieAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	movie := new(model.Movie)

	id := c.Params("id")
	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(id) {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid id format",
		})
	}
	result := db.Where(`id = ?`, id).First(&movie)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "movie not found",
		})
	}
	movie.MovieAPI = movieAPI
	db.Where(`id = ?`, id).Updates(&movie)

	message := fmt.Sprintf(`movie with id %s has been updated`, id)
	return c.JSON(fiber.Map{
		"message": message,
	})
}
