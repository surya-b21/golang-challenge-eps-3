package main

import (
	"challenge-eps3/app/routes"
	"challenge-eps3/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

// @title Golang Challenge eps 4
// @version 1.0.0
// @description golang challenge eps 4 api documentation
// @host localhost:8081
// @schemes http
// @BasePath /api/v1
func main() {
	services.InitDB()
	// db := services.DB
	// db.AutoMigrate(&model.Movie{})
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	routes.Handle(app)

	log.Fatal(app.Listen(":8081"))
}
