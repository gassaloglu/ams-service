package utils

import (
	"ams-service/internal/core/entities"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func BuildComparableQueryForField[T comparable](db *gorm.DB, cmp *entities.Comparable[T], field string) *gorm.DB {
	if (db == nil) || (cmp == nil) {
		return db
	}

	if cmp.EqualTo != nil {
		db = db.Where(field+" = ?", *cmp.EqualTo)
	}

	if cmp.NotEqaualTo != nil {
		db = db.Where(field+" != ?", *cmp.NotEqaualTo)
	}

	if cmp.GreaterThan != nil {
		db = db.Where(field+" > ?", *cmp.GreaterThan)
	}

	if cmp.LessThan != nil {
		db = db.Where(field+" < ?", *cmp.LessThan)
	}

	if cmp.GreaterThanOrEqualTo != nil {
		db = db.Where(field+" >= ?", *cmp.GreaterThanOrEqualTo)
	}
	if cmp.LessThanOrEqualTo != nil {
		db = db.Where(field+" <= ?", *cmp.LessThanOrEqualTo)
	}

	return db
}
