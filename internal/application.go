package internal

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/adapters/primary/rest/routes"
	postgresAdapter "ams-service/internal/adapters/secondary/postgres"
	"ams-service/internal/config"
	"ams-service/internal/core/services"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	setupLogger()

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

	// Load secrets
	scfg, err := config.LoadSecretConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load secrets")
	}

	// Initialize database connection based on configuration
	db, err := setupDatabaseConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}

	// Initialize repositories
	userRepo := postgresAdapter.NewUserRepositoryImpl(db)
	passengerRepo := postgresAdapter.NewPassengerRepositoryImpl(db)
	planeRepo := postgresAdapter.NewPlaneRepositoryImpl(db)
	flightRepo := postgresAdapter.NewFlightRepositoryImpl(db)
	employeeRepo := postgresAdapter.NewEmployeeRepositoryImpl(db)
	bankRepo := postgresAdapter.NewBankRepositoryImpl(db)

	// Initialize services
	bankService := services.NewBankService(bankRepo)
	tokenService := services.NewTokenService(scfg.JWTSecretKey)
	userService := services.NewUserService(userRepo, tokenService)
	planeService := services.NewPlaneService(planeRepo)
	employeeService := services.NewEmployeeService(employeeRepo, tokenService)
	flightService := services.NewFlightService(flightRepo)
	passengerService := services.NewPassengerService(passengerRepo, bankService, flightService)

	// Initialize controllers
	passengerController := controllers.NewPassengerController(passengerService)
	userController := controllers.NewUserController(userService)
	planeController := controllers.NewPlaneController(planeService)
	employeeController := controllers.NewEmployeeController(employeeService)
	flightController := controllers.NewFlightController(flightService)
	bankController := controllers.NewBankController(bankService)

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "AMS Service",
		ErrorHandler: middlewares.ErrorHandler,
	})

	// Add middleware
	app.Use(middlewares.TokenServiceInjector(tokenService))
	app.Use(middlewares.Logger())
	app.Use(cors.New())

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

func setupLogger() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Caller().
		Logger()
}

func buildDsn(db *config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
}

func setupDatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	switch cfg.Database.Type {
	case "postgres":
		dsn := buildDsn(&cfg.Database)
		return gorm.Open(pg.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		return nil, errors.New("unsupported database type")
	}
}
