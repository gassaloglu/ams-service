package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"fmt"
	"strings"

	"github.com/necmettindev/randomstring"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type PassengerService struct {
	repo   secondary.PassengerRepository
	bank   primary.BankService
	flight primary.FlightService
}

func NewPassengerService(repo secondary.PassengerRepository, bank primary.BankService, flight primary.FlightService) primary.PassengerService {
	return &PassengerService{repo: repo, bank: bank, flight: flight}
}

func (s *PassengerService) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByID(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error getting passenger by ID")
		return entities.Passenger{}, err
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully retrieved passenger by ID")
	return passenger, nil
}

func (s *PassengerService) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByPNR(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error getting passenger by PNR")
		return entities.Passenger{}, err
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully retrieved passenger by PNR")
	return passenger, nil
}

func (s *PassengerService) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	err := s.repo.OnlineCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		return err
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully checked in passenger")
	return nil
}

func (s *PassengerService) GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error) {
	passengers, err := s.repo.GetPassengersBySpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting passengers by specific flight")
		return nil, err
	}
	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved passengers by specific flight")
	return passengers, nil
}

func (s *PassengerService) CreatePassenger(request *entities.CreatePassengerRequest) (*entities.Passenger, error) {
	existingPassenger, err := s.repo.FindPassengersMatchingAnyUniquePassengerInfo(&request.Passenger)

	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}

	if existingPassenger != nil {
		return nil, fmt.Errorf("passenger with the same information already exists")
	}

	flight, err := s.flight.FindByFlightNumber(request.Passenger.FlightNumber)

	if err != nil {
		return nil, err
	}

	amount, err := calculateTicketPrice(flight.Price, request.Passenger.FareType)

	if err != nil {
		return nil, err
	}

	transaction, err := s.bank.Pay(&entities.PaymentRequest{
		Amount:     amount,
		CreditCard: request.CreditCard,
	})

	if err != nil {
		return nil, err
	}

	mappedPassenger, err := mapCreatePassengerRequestToPassengerEntity(request)

	if err != nil {
		return nil, err
	}

	mappedPassenger.TransactionId = transaction.ID
	mappedPassenger.FlightId = flight.ID

	passenger, err := s.repo.CreatePassenger(&mappedPassenger)

	if err != nil {
		return nil, err
	}

	return passenger, nil
}

func (p PassengerService) CreateAllPassengers(request *[]entities.CreatePassengerRequest) error {
	for _, passengerRequest := range *request {
		_, err := p.CreatePassenger(&passengerRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PassengerService) GetAllPassengers() ([]entities.Passenger, error) {
	passengers, err := s.repo.GetAllPassengers()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving all passengers")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved all passengers")
	return passengers, nil
}

func (s *PassengerService) EmployeeCheckInPassenger(request entities.EmployeeCheckInRequest) (entities.Passenger, error) {
	passenger, err := s.repo.EmployeeCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).
			Str("national_id", request.NationalId).
			Str("destination_airport", request.DestinationAirport).
			Msg("Error checking in passenger")
		return entities.Passenger{}, err
	}
	log.Info().
		Str("national_id", request.NationalId).
		Str("destination_airport", request.DestinationAirport).
		Msg("Successfully checked in passenger")
	return passenger, nil
}

func (s *PassengerService) CancelPassenger(request entities.CancelPassengerRequest) error {
	err := s.repo.CancelPassenger(request)
	if err != nil {
		log.Error().Err(err).Uint("passenger_id", request.PassengerID).Msg("Error canceling passenger")
		return err
	}
	log.Info().Uint("passenger_id", request.PassengerID).Msg("Successfully canceled passenger")
	return nil
}

var priceCoefficients = map[string]float64{
	"essentials": 1.0,
	"advantage":  1.2,
	"comfort":    1.2 * 1.2,
}

var extraBaggageMap = map[string]int{
	"essentials": 0,
	"advantage":  10,
	"comfort":    30,
}

func calculateTicketPrice(basePrice float64, fareType string) (float64, error) {
	coefficient, exists := priceCoefficients[fareType]

	if !exists {
		return 0, fmt.Errorf("invalid fare type: %s", fareType)
	}

	return basePrice * coefficient, nil
}

func mapCreatePassengerRequestToPassengerEntity(request *entities.CreatePassengerRequest) (entities.Passenger, error) {
	pnr, err := randomstring.GenerateString(randomstring.GenerationOptions{
		Length: 6,
	})

	if err != nil {
		return entities.Passenger{}, fmt.Errorf("failed to generate PNR: %w", err)
	}

	return entities.Passenger{
		PnrNo:        strings.ToUpper(pnr),
		FareType:     request.Passenger.FareType,
		NationalId:   request.Passenger.NationalID,
		Name:         request.Passenger.Name,
		Surname:      request.Passenger.Surname,
		Email:        request.Passenger.Email,
		Phone:        request.Passenger.Phone,
		Gender:       request.Passenger.Gender,
		Disabled:     request.Passenger.Disabled,
		Seat:         request.Passenger.Seat,
		BirthDate:    request.Passenger.BirthDate.Time,
		Child:        request.Passenger.Child,
		ExtraBaggage: extraBaggageMap[request.Passenger.FareType],
	}, nil
}
