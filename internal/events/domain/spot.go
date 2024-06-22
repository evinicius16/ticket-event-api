package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	InvalidSpotNumberError          = errors.New("Invalid spot number")
	SpotNotFoundError               = errors.New("Spot not found")
	SpotAlreadyReservedError        = errors.New("Spot already reserved")
	SpotNameRequiredError           = errors.New("Spot name is required")
	SpotNameLeastTwoCharactersError = errors.New("Spot name must be at least 2 characters long")
	SpotNameStartWithLetterError    = errors.New("Spot name must start with a letter")
	SpotNameEndsWithNumberError     = errors.New("Spot name must end with a number ")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (spot *Spot) Validate() error {

	if spot.Name == "" {
		return SpotNameRequiredError
	}

	if len(spot.Name) < 2 {
		return SpotNameLeastTwoCharactersError
	}

	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return SpotNameStartWithLetterError
	}

	if spot.Name[1] < '0' || spot.Name[1] > '9' {
		return SpotNameEndsWithNumberError
	}

	return nil
}
