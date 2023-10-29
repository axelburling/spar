package parser

import (
	"encoding/xml"
)

type info struct {
	KundNrLeveransMottagare string
	KundNrSlutkund          string
	OrgNrSlutkund           string
	UppdragsId              string
	SlutAnvandarId          string
	Tidsstampel             string
}

type Request struct {
	PersonId string
	Info     info
}

func Parse(body []byte) (Request, error) {
	var s Envelope
	err := xml.Unmarshal(body, &s)
	if err != nil {
		return Request{}, err
	}

	personId := s.Body.SPARPersonsokningFraga.PersonsokningFraga.PersonId.FysiskPersonId

	info := info{
		KundNrLeveransMottagare: s.Body.SPARPersonsokningFraga.IdentifieringsInformation.KundNrLeveransMottagare,
		KundNrSlutkund:          s.Body.SPARPersonsokningFraga.IdentifieringsInformation.KundNrSlutkund,
		OrgNrSlutkund:           s.Body.SPARPersonsokningFraga.IdentifieringsInformation.OrgNrSlutkund,
		UppdragsId:              s.Body.SPARPersonsokningFraga.IdentifieringsInformation.UppdragsId,
		SlutAnvandarId:          s.Body.SPARPersonsokningFraga.IdentifieringsInformation.SlutAnvandarId,
		Tidsstampel:             s.Body.SPARPersonsokningFraga.IdentifieringsInformation.Tidsstampel,
	}

	req := Request{
		PersonId: personId,
		Info:     info,
	}

	return req, nil
}
