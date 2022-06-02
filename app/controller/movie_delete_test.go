package controller

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteMovie(t *testing.T) {
	app := fiber.New()
	app.Delete("/movies/:id", DeleteMovie)

	req, _ := http.NewRequest("DELETE", "/movies/1", nil)
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	req, _ = http.NewRequest("DELETE", "/movies/satu", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, res.StatusCode, "response code")

	req, _ = http.NewRequest("DELETE", "/movies/2", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, res.StatusCode, "response code")
}
