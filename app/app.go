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
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var LOG_PREFIX string = "app.go"

func Run() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("%s - Failed to load configuration: %v", LOG_PREFIX, err)
	}

	var userRepo ports.UserRepository
	var passengerRepo ports.PassengerRepository
	var planeRepo ports.PlaneRepository
	var flightRepo ports.FlightRepository

	// Initialize database connection based on configuration
	switch cfg.Database.Type {
	case "postgres":
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name))
		if err != nil {
			log.Fatalf("%s - Failed to connect to PostgreSQL database: %v", LOG_PREFIX, err)
		}
		defer db.Close()
		userRepo = postgres.NewUserRepositoryImpl(db)
		passengerRepo = postgres.NewPassengerRepositoryImpl(db)
		planeRepo = postgres.NewPlaneRepositoryImpl(db)
		flightRepo = postgres.NewFlightRepositoryImpl(db)
	default:
		log.Fatalf("%s - Unsupported database type: %s", LOG_PREFIX, cfg.Database.Type)
	}

	// Initialize services
	passengerService := services.NewPassengerService(passengerRepo)
	userService := services.NewUserService(userRepo)
	planeService := services.NewPlaneService(planeRepo)
	flightService := services.NewFlightService(flightRepo)


	// Initialize controllers
	passengerController := controllers.NewPassengerController(passengerService)
	userController := controllers.NewUserController(userService)
	planeController := controllers.NewPlaneController(planeService)
	flightController := controllers.NewFlightController(flightService)

	// Setup router
	router := gin.Default()
	router.Use(middlewares.Logger())
	router.Use(middlewares.ErrorHandler())

	// Register routes
	RegisterPlaneRoutes(router, planeController)
	RegisterFlightRoutes(router, flightController)
	RegisterPassengerRoutes(router, passengerController)
	RegisterUserRoutes(router, userController)

	planeRoute := router.Group("/plane")
	{
		planeRoute.GET("/all", planeController.GetAllPlanes)
		planeRoute.POST("/add", planeController.AddPlane)
		planeRoute.PUT("/status", planeController.SetPlaneStatus)
		planeRoute.GET("/registration", planeController.GetPlaneByRegistration)
		planeRoute.GET("/flightnumber", planeController.GetPlaneByFlightNumber)
		planeRoute.GET("/location", planeController.GetPlaneByLocation)
	}

	// Run the server
	err = router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Failed to start server: %v", LOG_PREFIX, err))
	}
}
