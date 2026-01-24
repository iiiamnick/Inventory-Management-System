package handlers

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/iiiamnick/Inventory-Management-System/config"
)

func BenchmarkCarGet(b *testing.B) {
	config.ConnectDB()
	app := fiber.New()
	app.Get("/cars/:id", GetCar)

	req, _ := http.NewRequest("GET", "/cars/5", nil)
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {

		_, _ = app.Test(req, 5000)
	}

}
