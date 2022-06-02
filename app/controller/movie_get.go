package controller

import (
	"challenge-eps3/app/model"
	"challenge-eps3/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetMovie function
func GetMovie(c *fiber.Ctx) error {
	db := services.DB
	var movies []model.Movie
	db.Model(&model.Movie{}).Find(&movies)
	return c.JSON(movies)
}
