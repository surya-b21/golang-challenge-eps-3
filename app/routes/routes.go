package routes

import (
	"challenge-eps3/app/controller"

	"github.com/gofiber/fiber/v2"
)

// Handle route function
func Handle(app *fiber.App) {
	// app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/movies", controller.GetMovie)
	api.Post("/movies", controller.PostMovie)
	api.Put("/movies/:id", controller.PutMovie)
	api.Get("/movies/:id", controller.GetMovieID)
	api.Delete("/movies/:id", controller.DeleteMovie)
}
