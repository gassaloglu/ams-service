package entities

type FraudulentActivity struct {
	Transaction Transaction `json:"transaction"`
	Passenger   Passenger   `json:"passenger"`
	User        User        `json:"user"`
}
