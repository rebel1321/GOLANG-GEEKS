package main

import (
	"car/config"
	"car/handlers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "car/docs"
)

// @title Car Inventory
// @version 1.0
// @description This is an API spec for car inventory management
// @host localhost:8081
// @BasePath /
func main() {
	config.ConnectDB()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(logger.New())

	// app.Use(middlewares.SecurityHeaders)

	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "12345",
	// 		"manager": "67890",
	// 		"employee": "abcde",
	// 	},
	// 	Unauthorized: func(c *fiber.Ctx) error {
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 			"error": "Unauthorized",
	// 		})
	// 	},
	// }))

	app.Use(etag.New())
	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Put("/cars/:id", handlers.UpdateCar)

	fmt.Println("Fiber Server started at :8081")
	if err := app.Listen(":8081"); err != nil {
		log.Fatal(err)
	}
}
