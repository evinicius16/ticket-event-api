package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{partnerBaseURLs: partnerBaseURLs}
}

func (defaultPartner *DefaultPartnerFactory) CreatePartner(partnerID int) (Partner, error) {
	baseURL, ok := defaultPartner.partnerBaseURLs[partnerID]
	if !ok {
		return nil, fmt.Errorf("partner with ID %d not fount", partnerID)
	}

	switch partnerID {
	case 1:
		return &PartnerOne{BaseUrl: baseURL}, nil
	default:
		return nil, fmt.Errorf("partner with ID %d not found", partnerID)
	}

}
