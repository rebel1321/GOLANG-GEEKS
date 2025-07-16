package handlers

import (
	"car/models"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)


var mu sync.Mutex


func CreateCar(c *fiber.Ctx) error{
	mu.Lock()
	defer mu.Unlock()

	car := models.Car{}
	if err := c.BodyParser(&car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"details": err.Error(),
		})
	}

	car.Insert()
	fmt.Printf("Car created: %+v\n", car)
	return c.Status(fiber.StatusCreated).JSON(car)
}
func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car ID",
			"details": err.Error(),
		})
	}

	car := &models.Car{Id: id}
	if err := car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car ID",
			"details": err.Error(),
		})
	}

	car := &models.Car{Id: id}
	if err := car.Delete(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found",
			"details": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car ID",
			"details": err.Error(),
		})
	}

	car := &models.Car{Id: id}
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"details": err.Error(),
		})
	}

	car.Update()
	fmt.Printf("Car updated: %+v\n", car)
	return c.Status(fiber.StatusOK).JSON(car)
}
