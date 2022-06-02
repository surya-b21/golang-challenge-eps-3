package controller

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetMovieID(t *testing.T) {
	app := fiber.New()

	app.Get("/movies/:id", GetMovieID)

	req, _ := http.NewRequest("GET", "/movies/1", nil)
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	req, _ = http.NewRequest("GET", "/movies/2", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, res.StatusCode, "response code")

	req, _ = http.NewRequest("GET", "/movies/satu", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, res.StatusCode, "response code")
}
