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
	config.ConnectMongo()
	config.ConnectCache()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(logger.New())

	

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
