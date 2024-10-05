package middleware

import (
	"UserService/internal/models" // Adjust the import path as necessary

	"github.com/gofiber/fiber/v2"
)

// ResponseMiddleware is a middleware that formats the response
func ResponseMiddleware(c *fiber.Ctx) error {
	// Call the next handler
	err := c.Next()

	// If there was an error, handle it
	if err != nil {
		return err
	}

	// Create a standardized response
	response := models.ResponseFormat{
		Status:  "success",
		Message: "Request was successful",
		Data:    c.Locals("data"), // Use Locals to get data set by handlers
	}

	// If the status code is an error (e.g., 4xx, 5xx), change the format accordingly
	if c.Response().StatusCode() >= 400 {
		response.Status = "error"
		response.Message = "An error occurred"
		response.Data = nil
	}
	// Set the response header and return the JSON response
	return c.JSON(response)
}
