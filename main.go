package main

import (
	"challenge-eps3/app/routes"
	"challenge-eps3/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

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
