package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type FlightController struct {
	service primary.FlightService
}

func NewFlightController(service primary.FlightService) *FlightController {
	return &FlightController{service: service}
}

func (c *FlightController) GetSpecificFlight(ctx *fiber.Ctx) error {
	var request entities.GetSpecificFlightRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	userOrEmployeeID, err := utils.ExtractUserOrEmployeeID(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Unauthorized access")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	log.Info().Str("id", userOrEmployeeID).Msg("User/Employee attempting to get a specific flight")
	flight, err := c.service.GetSpecificFlight(request, userOrEmployeeID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Flight not found"})
		}
		log.Error().Err(err).Msg("Error getting specific flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flight"})
	}

	return ctx.Status(http.StatusOK).JSON(flight)
}

func (c *FlightController) GetAllFlights(ctx *fiber.Ctx) error {
	allowedRoles := []string{"admin", "flight_planner", "passenger_services", "ground_services"}
	role, err := utils.CheckRoleAuthorization(ctx, allowedRoles)
	if err != nil {
		log.Error().Err(err).Msg("Unauthorized access")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	log.Info().Str("role", role).Msg("Authorized role attempting to view all flights")
	flights, err := c.service.GetAllFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllFlightsDestinationDateFlights(ctx *fiber.Ctx) error {
	userOrEmployeeID, err := utils.ExtractUserOrEmployeeID(ctx)
	if err != nil {
		return err
	}

	log.Info().Str("id", userOrEmployeeID).Msg("User/Employee attempting to view specific flights")
	var request entities.GetAllFlightsDestinationDateRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	flights, err := c.service.GetAllFlightsDestinationDateFlights(request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting specific flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllActiveFlights(ctx *fiber.Ctx) error {
	allowedRoles := []string{"admin", "flight_planner", "passenger_services", "ground_services"}
	role, err := utils.CheckRoleAuthorization(ctx, allowedRoles)
	if err != nil {
		log.Error().Err(err).Msg("Unauthorized access")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	log.Info().Str("role", role).Msg("Authorized role attempting to view all active flights")
	flights, err := c.service.GetAllActiveFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all active flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all active flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No active flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "There are no active flights"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) CancelFlight(ctx *fiber.Ctx) error {
	allowedRoles := []string{"admin", "flight_planner"}
	role, err := utils.CheckRoleAuthorization(ctx, allowedRoles)
	if err != nil {
		log.Error().Err(err).Msg("Unauthorized access")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	log.Info().Str("role", role).Msg("Authorized role attempting to cancel a flight")
	var request entities.CancelFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = c.service.CancelFlight(request)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found to cancel")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flight found to cancel"})
		}
		log.Error().Err(err).Msg("Error canceling flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error canceling flight"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Flight canceled successfully"})
}

func (c *FlightController) AddFlight(ctx *fiber.Ctx) error {
	// Check if the user is authorized with the required roles
	allowedRoles := []string{"admin", "flight_planner"}
	role, err := utils.CheckRoleAuthorization(ctx, allowedRoles)
	if err != nil {
		log.Error().Err(err).Msg("Unauthorized access")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	log.Info().Str("role", role).Msg("Authorized role attempting to add a flight")

	// Parse the request body
	var request entities.AddFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Call the service to add the flight
	err = c.service.AddFlight(request)
	if err != nil {
		if err.Error() == "pq: insert or update on table \"flights\" violates foreign key constraint \"fk_plane_registration\"" {
			log.Error().Err(err).Msg("Foreign key constraint violation")
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid plane registration"})
		}
		log.Error().Err(err).Msg("Error adding flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error adding flight"})
	}

	log.Info().Msg("Successfully added flight")
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Flight added successfully"})
}
