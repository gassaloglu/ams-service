package app

import (
	"ams-service/application/ports"
	"ams-service/config"
	"ams-service/core/services"
	"ams-service/infrastructure/api/controllers"
	"ams-service/infrastructure/persistence/repositories/postgres"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var LOG_PREFIX string = "app.go"

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

	var userRepo ports.UserRepository
	var passengerRepo ports.PassengerRepository
	var planeRepo ports.PlaneRepository
	var employeeRepo ports.EmployeeRepository
	var flightRepo ports.FlightRepository
	var bankRepo ports.BankRepository

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

	// Setup router
	gin.SetMode("release")
	router := gin.Default()
	router.Use(middlewares.Logger())
	router.Use(middlewares.ErrorHandler())

	// Register routes
	RegisterPlaneRoutes(router, planeController)
	RegisterFlightRoutes(router, flightController)
	RegisterPassengerRoutes(router, passengerController)
	RegisterUserRoutes(router, userController)
	RegisterEmployeeRoutes(router, employeeController)
	RegisterBankRoutes(router, bankController)

	// Run the server
	log.Info().
		Str("port", cfg.ServerPort).
		Msg("Starting REST server")

	err = router.Run(fmt.Sprintf(":%s", cfg.ServerPort))

	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
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
