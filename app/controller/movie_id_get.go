package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// GetMovieID function
// @Summary     get a movie
// @Description get a movie data
// @Tags         movies
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  model.Movie
// @Failure      400
// @Failure      404
// @Router       /movies/{id} [get]
func GetMovieID(c *fiber.Ctx) error {
	id := c.Params("id")
	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(id) {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid id format",
		})
	}
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
