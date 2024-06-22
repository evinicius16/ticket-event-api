package domain

import (
	"errors"
	"time"
)

var (
	ErrorEventNameRequired = errors.New("event name is required")
	ErrorEvendDateFuture   = errors.New("event date must be in the future")
	ErrorEventCapacityZero = errors.New("event capacity must be greater than zero")
	ErrorEventPriceZero    = errors.New("event price must be grater than zero")
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
		return ErrorEventNameRequired
	}

	if event.Date.Before(time.Now()) {
		return ErrorEvendDateFuture
	}

	if event.Capacity <= 0 {
		return ErrorEventCapacityZero
	}

	if event.Price <= 0 {
		return ErrorEventPriceZero
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
