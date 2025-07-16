package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Request entered | PATH: %v | METHOD: %v | TIME: %v\n", r.URL.Path, r.Method, start)
		next.ServeHTTP(w, r)

		fmt.Println("Request completed | PATH:", r.URL.Path, "| METHOD:", r.Method, "| DURATION:", time.Since(start))
	})
}

func SecurityHeaders(c *fiber.Ctx) error{
	c.Response().Header.Add("X-Content-Type-Options", "nosniff")
	c.Response().Header.Add("X-XSS-Protection", "1; mode=block")
	c.Response().Header.Add("Strict-Transport-Security", "max-age=31536000 includeSubDomains")
	c.Response().Header.Add("Content-Security-Policy", "default-src 'self'")
	
	return c.Next()
	
}