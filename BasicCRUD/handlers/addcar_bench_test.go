package handlers

import (
	"car/config"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func BenchmarkCarGet(b *testing.B) {
	// Test implementation goes here
	config.ConnectDB()
	app := fiber.New()
	// Define routes
	app.Get("/cars/:id", GetCar)

	req, _ := http.NewRequest("GET", "/cars/5", nil)
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req, 5000)
	}

	
}