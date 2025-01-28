package app

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/infrastructure/api/controllers"
	"ams-service/middlewares"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	// Mock repositories
	userRepo := &MockUserRepository{}
	passengerRepo := &MockPassengerRepository{}

	// Initialize services
	passengerService := services.NewPassengerService(passengerRepo)
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	passengerController := controllers.NewPassengerController(passengerService)
	userController := controllers.NewUserController(userService)

	// Setup router
	router := gin.Default()
	router.Use(middlewares.Logger())
	router.Use(middlewares.ErrorHandler())

	// Setup routes
	passengerRoute := router.Group("/passenger")
	{
		passengerRoute.POST("/checkin", passengerController.OnlineCheckInPassenger)
		passengerRoute.GET("/:id", passengerController.GetPassengerByID)
	}

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", userController.RegisterUser)
	}

	return router
}

func TestRegisterUser(t *testing.T) {
	router := setupRouter()

	user := entities.User{
		Name:         "John",
		Surname:      "Doe",
		Username:     "johndoe",
		Email:        "john.doe@example.com",
		PasswordHash: "hashedpassword",
		Phone:        "1234567890",
		Gender:       "male",
		BirthDate:    time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["message"] != "Registration successful" {
		t.Fatalf("Expected message 'Registration successful', but got '%s'", response["message"])
	}
}

func TestOnlineCheckInPassenger(t *testing.T) {
	router := setupRouter()

	checkInRequest := entities.OnlineCheckInRequest{
		PNR:     "ABC123",
		Surname: "Doe",
	}

	jsonValue, _ := json.Marshal(checkInRequest)
	req, _ := http.NewRequest("POST", "/passenger/checkin", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["message"] != "Check-in successful" {
		t.Fatalf("Expected message 'Check-in successful', but got '%s'", response["message"])
	}
}

func TestGetPassengerByID(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/passenger/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var passenger entities.Passenger
	err := json.Unmarshal(w.Body.Bytes(), &passenger)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if passenger.NationalId != "12345678901" {
		t.Fatalf("Expected NationalId '12345678901', but got '%s'", passenger.NationalId)
	}
}
