package domain

import "errors"

type TicketType string

var ErrorTicketPriceZero = errors.New("ticket price must be greater than zero")

const (
	TickeTypeHalf  TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TickeTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatePrice() {

	if ticket.TicketType == TickeTypeHalf {
		ticket.Price /= 2
	}
}

func (ticket *Ticket) Validate() error {
	if ticket.Price <= 0 {
		return ErrorTicketPriceZero
	}

	return nil
}
