package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventById(eventID string) (*Event, error)
	FindSpotsByEventId(eventID string) ([]*Spot, error)
	findSpotByName(eventID, spotName string) (*Spot, error)
	// CreateEvent(event *Event) error
	// CreateSpot(spot *Spot) error
	// CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}
