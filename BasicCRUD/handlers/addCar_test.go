package handlers

import (
	"car/config"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCarAdd(t *testing.T) {
	// Test implementation goes here
	config.ConnectDB()
	app:=fiber.New()
	// Define routes
	app.Post("/cars", CreateCar)

	body :=`{
		"name": "Test Car",
		"model": "Model X",
		"brand": "Brand Y",
		"year": 2022,
		"price": 500000.00
	}`
	req,_ :=http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 5000)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t,fiber.StatusCreated, resp.StatusCode)
}
func TestCarGet(t *testing.T) {
	// Test implementation goes here
	config.ConnectDB()
	app:=fiber.New()
	// Define routes
	app.Get("/cars/:id", GetCar)


	req,_ :=http.NewRequest("GET", "/cars/5", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 5000)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t,fiber.StatusOK, resp.StatusCode)
}