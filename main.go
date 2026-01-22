package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iiiamnick/Inventory-Management-System/config"
	"github.com/iiiamnick/Inventory-Management-System/handlers"
	"github.com/iiiamnick/Inventory-Management-System/middlewares"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(middlewares.SecurityHeaders)
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin":   "123",
			"manager": "quert",
			"john":    "wick",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "user is not authorised",
			})

		},
	}))

	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Get("/cars/:id", handlers.DeleteCar)

	fmt.Println("Server is Listening at port :3030")
	err := app.Listen(":3030")
	if err != nil {
		log.Fatalln("Could not listen at the port", err)
	}

}
