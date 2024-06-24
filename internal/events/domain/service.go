package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

var (
	ErrorInvalidQuantity = errors.New("quantity must be greater then zero")
)

func NewSpotService() *spotService {
	return &spotService{}
}

func (spotService *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrorInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)

		spot, err := NewSpot(event, spotName)

		if err != nil {
			return err
		}

		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
