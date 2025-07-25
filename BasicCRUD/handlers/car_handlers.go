package handlers

import (
	"car/models"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)


var mu sync.Mutex

// CreateCar godoc
// @Summary      Create a new car
// @Description  Adds a new car to inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        car  body      models.Car  true  "Car to add"
// @Success      201  {object}  models.Car
// @Failure      400  {object}  map[string]string
// @Router       /cars [post]

// CreateCar handles the HTTP POST request to create a new car.
// It parses the request body into a Car model, inserts the car into the database,
// and returns the created car as a JSON response with status 201 Created.
// If the request body is invalid, it returns a 400 Bad Request error with details.
func CreateCar(c *fiber.Ctx) error {
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

// GetCar godoc
// @Summary      Get a car
// @Description  Get a car from inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id  path      string true   "Car ID"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /cars/{id} [get]
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

// DeleteCar godoc
// @Summary      Delete a car
// @Description  Delete a car from inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id  path      string true   "Car ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /cars/{id} [delete]
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

// UpdateCar godoc
// @Summary      Update a car
// @Description  Update a car in inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id  path      string true   "Car ID"
// @Param        car  body      models.Car  true  "Car to update"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /cars/{id} [put]
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
