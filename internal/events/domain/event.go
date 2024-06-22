package domain

import (
	"errors"
	"time"
)

var (
	EventNameRequiredError = errors.New("Event name is required")
	EvendDateFutureError   = errors.New("Event date must be in the future")
	EventCapacityZeroError = errors.New("Event capacity must be greater than zero")
	EventPriceZeroError    = errors.New("Event price must be grater than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURl     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (event Event) Validate() error {
	if event.Name == "" {
		return EventNameRequiredError
	}

	if event.Date.Before(time.Now()) {
		return EvendDateFutureError
	}

	if event.Capacity <= 0 {
		return EventCapacityZeroError
	}

	if event.Price <= 0 {
		return EventPriceZeroError
	}

	return nil
}

func (event *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(event, name)

	if err != nil {
		return nil, err
	}

	event.Spots = append(event.Spots, *spot)

	return spot, nil
}
