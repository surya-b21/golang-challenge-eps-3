package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetMovie function
// @Summary      get all movies
// @Description  get all movies data
// @Tags         movies
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  model.Movie
// @Failure      400
// @Router       /movies [get]
func GetMovie(c *fiber.Ctx) error {
	db := services.DB
	var movies []model.Movie
	db.Model(&model.Movie{}).Find(&movies)
	return c.JSON(movies)
}
