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
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var LOG_PREFIX string = "app.go"

func Run() {
	// Load default environment variables from .env file first
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s - Failed to load .env file: %v", LOG_PREFIX, err)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("%s - Failed to load configuration: %v", LOG_PREFIX, err)
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
			middlewares.LogError(fmt.Sprintf("%s - Failed to connect to PostgreSQL database: %v", LOG_PREFIX, err))
			return
		}
		defer db.Close()

		// Ping the database to test the connection
		if err := db.Ping(); err != nil {
			middlewares.LogError(fmt.Sprintf("%s - Failed to ping PostgreSQL database: %v", LOG_PREFIX, err))
			return
		}

		middlewares.LogInfo(fmt.Sprintf("%s - Connected to PostgreSQL database", LOG_PREFIX))

		userRepo = postgres.NewUserRepositoryImpl(db)
		passengerRepo = postgres.NewPassengerRepositoryImpl(db)
		planeRepo = postgres.NewPlaneRepositoryImpl(db)
		flightRepo = postgres.NewFlightRepositoryImpl(db)
		employeeRepo = postgres.NewEmployeeRepositoryImpl(db)
		bankRepo = postgres.NewBankRepositoryImpl(db)
	default:
		log.Fatalf("%s - Unsupported database type: %s", LOG_PREFIX, cfg.Database.Type)
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
	err = router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Failed to start server: %v", LOG_PREFIX, err))
	}
}
