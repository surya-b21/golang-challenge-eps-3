package controller

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutMovie(t *testing.T) {
	app := fiber.New()
	app.Put("/movies/:id", PutMovie)
	payload := bytes.NewReader([]byte(`
	{
		"title": "Boy Story",
		"year": 1999,
		"summary": "Good Strong",
		"director": "pak sutradara"
	}
	`))

	req, _ := http.NewRequest("PUT", "/movies/1", payload)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	req, _ = http.NewRequest("PUT", "/movies/satu", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, res.StatusCode, "response code")

	req, _ = http.NewRequest("PUT", "/movies/2", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, res.StatusCode, "response code")
}
