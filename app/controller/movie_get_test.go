package controller

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetMovie(t *testing.T) {
	app := fiber.New()
	app.Get("/movies", GetMovie)

	req, _ := http.NewRequest("GET", "/movies", nil)
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get response")
}
