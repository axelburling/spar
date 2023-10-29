package parser

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                   string `xml:",chardata"`
		SPARPersonsokningFraga struct {
			Text                      string `xml:",chardata"`
			Xmlns                     string `xml:"xmlns,attr"`
			ID                        string `xml:"id,attr"`
			Sok                       string `xml:"sok,attr"`
			Argument                  string `xml:"argument,attr"`
			Person                    string `xml:"person,attr"`
			IdentifieringsInformation struct {
				Text                    string `xml:",chardata"`
				KundNrLeveransMottagare string `xml:"KundNrLeveransMottagare"`
				KundNrSlutkund          string `xml:"KundNrSlutkund"`
				OrgNrSlutkund           string `xml:"OrgNrSlutkund"`
				UppdragsId              string `xml:"UppdragsId"`
				SlutAnvandarId          string `xml:"SlutAnvandarId"`
				Tidsstampel             string `xml:"Tidsstampel"`
			} `xml:"IdentifieringsInformation"`
			PersonsokningFraga struct {
				Text     string `xml:",chardata"`
				PersonId struct {
					Text           string `xml:",chardata"`
					FysiskPersonId string `xml:"FysiskPersonId"`
				} `xml:"PersonId"`
			} `xml:"PersonsokningFraga"`
		} `xml:"SPARPersonsokningFraga"`
	} `xml:"Body"`
}
