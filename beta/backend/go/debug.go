// beta/backend/go/debug.go

package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// DebugMiddleware is a middleware that logs request and response details
func DebugMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()

	// Log request details
	log.Printf("Incoming request: %s %s", c.Method(), c.Path())

	// Process the request
	err := c.Next()

	// Log response details
	duration := time.Since(startTime)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Response status: %d, Duration: %s", c.Response().StatusCode(), duration)
	}

	return err
}

// LogError logs error messages
func LogError(err error) {
	log.Printf("Error: %v", err)
}

// LogInfo logs informational messages
func LogInfo(message string) {
	log.Printf("Info: %s", message)
}
