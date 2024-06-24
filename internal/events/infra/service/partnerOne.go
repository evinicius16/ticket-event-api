package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PartnerOne struct {
	BaseUrl string
}

type PartnerOneReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	Email      string   `json:"email"`
}

type PartnerOneReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

func (partner *PartnerOne) MakeReservation(request *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := PartnerOneReservationRequest{
		Spots:      request.Spots,
		TicketKind: request.TicketType,
		Email:      request.Email,
	}

	body, err := json.Marshal(partnerRequest)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", partner.BaseUrl, request.EventID)

	httpRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}

	var partnerResponse []PartnerOneReservationResponse

	if err := json.NewDecoder(httpResponse.Body).Decode(&partnerResponse); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResponse))
	for i, r := range partnerResponse {
		responses[i] = ReservationResponse{
			ID:     r.ID,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}

	return responses, nil
}
