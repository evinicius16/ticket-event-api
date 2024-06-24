package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrorInvalidSpotNumber          = errors.New("invalid spot number")
	ErrorSpotNotFound               = errors.New("spot not found")
	ErrorSpotAlreadyReserved        = errors.New("spot already reserved")
	ErrorSpotNameRequired           = errors.New("spot name is required")
	ErrorSpotNameLeastTwoCharacters = errors.New("spot name must be at least 2 characters long")
	ErrorSpotNameStartWithLetter    = errors.New("spot name must start with a letter")
	ErrorSpotNameEndsWithNumber     = errors.New("spot name must end with a number ")
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
		return ErrorEventNameRequired
	}

	if len(spot.Name) < 2 {
		return ErrorSpotNameLeastTwoCharacters
	}

	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return ErrorSpotNameStartWithLetter
	}

	if spot.Name[1] < '0' || spot.Name[1] > '9' {
		return ErrorSpotNameEndsWithNumber
	}

	return nil
}

func (spot *Spot) Reserve(ticketID string) error {
	if spot.Status == SpotStatusSold {
		return ErrorSpotAlreadyReserved
	}

	spot.Status = SpotStatusSold
	spot.TicketID = ticketID

	return nil
}
