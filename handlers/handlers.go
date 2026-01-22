package handlers

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/iiiamnick/Inventory-Management-System/models"
)

var mu sync.Mutex

func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}
	err := c.BodyParser(car)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Incorrect input body", "details": err.Error(),
		})
	}
	car.Insert()
	fmt.Println("Car saved to the inventory with the given id :", car.ID)
	return c.Status(fiber.StatusCreated).JSON(car)
}

func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	car.ID = id
	err = car.GET()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "car with given id not found",
			"id":    car.ID,
		})
	}

	fmt.Println("Car with the given id found:", id)
	return c.Status(fiber.StatusOK).JSON(car)

}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id", "details": err.Error(),
		})
	}
	car.ID = id
	car.DELETE()

	fmt.Println("Deletion is completed for the given id", id)
	return c.SendStatus(fiber.StatusNoContent)

}
