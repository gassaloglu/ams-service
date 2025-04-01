package entities

type Plane struct {
	PlaneRegistration    string `json:"plane_registration" gorm:"primaryKey;size:10;not null"`
	PlaneType            string `json:"plane_type" gorm:"size:50;not null"`
	Location             string `json:"location" gorm:"size:50;not null"`
	TotalPassengers      int    `json:"total_passengers" gorm:"not null"`
	MaxPassengers        int    `json:"max_passengers" gorm:"not null"`
	EconomyPassengers    int    `json:"economy_passengers" gorm:"not null"`
	BusinessPassengers   int    `json:"business_passengers" gorm:"not null"`
	FirstClassPassengers int    `json:"first_class_passengers" gorm:"not null"`
	FlightNumber         string `json:"flight_number" gorm:"size:10;not null"`
	IsAvailable          bool   `json:"is_available" gorm:"not null"`
}
