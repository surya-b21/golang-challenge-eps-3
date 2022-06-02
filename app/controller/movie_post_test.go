package controller

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// TestPostMovie function
func TestPostMovie(t *testing.T) {
	app := fiber.New()
	app.Post("/movies", PostMovie)
	payload := bytes.NewReader([]byte(`
	{
		"title": "Boy Story",
		"year": 1999,
		"summary": "Good Strong",
		"director": "pak sutradara"
	}
	`))

	req, _ := http.NewRequest("POST", "/movies", payload)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 201, res.StatusCode, "response code")

	req, _ = http.NewRequest("POST", "/movies", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, res.StatusCode, "response code")

}
