package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// IsBatchRequest checks if the incoming request is a batch request.
// It determines this by looking for a query parameter named "batch"
// and checking if its value is "true" (case-insensitive).
//
// Parameters:
//
//	c (*fiber.Ctx): The Fiber context associated with the request.
//
// Returns:
//
//	bool: True if the "batch" query parameter is present and set to "true",
//	      false otherwise.
func IsBatchRequest(c *fiber.Ctx) bool {
	// Get the value of the "batch" query parameter.
	// If the parameter is not present, Query returns an empty string.
	batchParam := c.Query("batch")

	// Convert the parameter value to lowercase for case-insensitive comparison.
	isBatch := strings.ToLower(batchParam) == "true"

	// Return the result.
	return isBatch
}
