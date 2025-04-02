package internal

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/adapters/primary/rest/routes"
	"ams-service/internal/adapters/secondary/postgres"
	"ams-service/internal/config"
	"ams-service/internal/core/services"
	"ams-service/internal/ports/secondary"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() {
	initLogger()

	// Load default environment variables from .env file first
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env file")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	var userRepo secondary.UserRepository
	var passengerRepo secondary.PassengerRepository
	var planeRepo secondary.PlaneRepository
	var employeeRepo secondary.EmployeeRepository
	var flightRepo secondary.FlightRepository
	var bankRepo secondary.BankRepository

	// Initialize database connection based on configuration
	switch cfg.Database.Type {
	case "postgres":
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SSLMode))

		if err != nil {
			log.Error().Err(err).Msg("Failed to connect to PostgreSQL database")
			return
		}

		defer db.Close()

		// Ping the database to test the connection
		if err := db.Ping(); err != nil {
			log.Error().Err(err).Msg("Failed to ping PostgreSQL database")
			return
		}

		log.Info().Msg("Connected to PostgreSQL database")

		userRepo = postgres.NewUserRepositoryImpl(db)
		passengerRepo = postgres.NewPassengerRepositoryImpl(db)
		planeRepo = postgres.NewPlaneRepositoryImpl(db)
		flightRepo = postgres.NewFlightRepositoryImpl(db)
		employeeRepo = postgres.NewEmployeeRepositoryImpl(db)
		bankRepo = postgres.NewBankRepositoryImpl(db)
	default:
		log.Fatal().Msgf("Unsupported database type %s", cfg.Database.Type)
	}

	// Initialize services
	passengerService := services.NewPassengerService(passengerRepo)
	userService := services.NewUserService(userRepo)
	planeService := services.NewPlaneService(planeRepo)
	employeeService := services.NewEmployeeService(employeeRepo)
	flightService := services.NewFlightService(flightRepo)
	bankService := services.NewBankService(bankRepo)

	// Initialize controllers
	passengerController := controllers.NewPassengerController(passengerService)
	userController := controllers.NewUserController(userService)
	planeController := controllers.NewPlaneController(planeService)
	employeeController := controllers.NewEmployeeController(employeeService)
	flightController := controllers.NewFlightController(flightService)
	bankController := controllers.NewBankController(bankService)

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		AppName: "AMS Service",
	})

	// Add middleware
	app.Use(middlewares.Logger())
	app.Use(middlewares.ErrorHandler())

	// Register routes
	routes.RegisterUserRoutes(app, userController)
	routes.RegisterEmployeeRoutes(app, employeeController)
	routes.RegisterPassengerRoutes(app, passengerController)
	routes.RegisterPlaneRoutes(app, planeController)
	routes.RegisterFlightRoutes(app, flightController)
	routes.RegisterBankRoutes(app, bankController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info().Str("port", port).Msg("Starting server")
	if err := app.Listen(":" + port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func initLogger() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Caller().
		Logger()
}
