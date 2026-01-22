package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/iiiamnick/Inventory-Management-System/config"
	"github.com/stretchr/testify/assert"
)

func TestcarAdd(t *testing.T) {
	config.ConnectDB()
	app := fiber.New()
	app.Post("/cars", CreateCar)

	body := `{
		"name": "sample book",
    "model": "nick",
    "brand": "audi",
    "year": 2024,
    "price": 100
	}`

	req, _ := http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
