package formater

type Response struct {
	SEnvelope SEnvelope `xml:"S:Envelope"`
}
type PersonID struct {
	FysiskPersonID int `xml:"FysiskPersonId"`
}
type Persondetaljer struct {
	DatumFrom            string `xml:"DatumFrom"`
	DatumTill            string `xml:"DatumTill"`
	Fornamn              string `xml:"Fornamn"`
	Tilltalsnamn         string `xml:"Tilltalsnamn"`
	Efternamn            string `xml:"Efternamn"`
	Sekretessmarkering   string `xml:"Sekretessmarkering"`
	SkyddadFolkbokforing string `xml:"SkyddadFolkbokforing"`
	Fodelsetid           string `xml:"Fodelsetid"`
	Kon                  string `xml:"Kon"`
}
type Folkbokforingsadress struct {
	DatumFrom                string `xml:"DatumFrom"`
	DatumTill                string `xml:"DatumTill"`
	Utdelningsadress2        string `xml:"Utdelningsadress2"`
	PostNr                   int    `xml:"PostNr"`
	Postort                  string `xml:"Postort"`
	FolkbokfordLanKod        int    `xml:"FolkbokfordLanKod"`
	FolkbokfordKommunKod     int    `xml:"FolkbokfordKommunKod"`
	Folkbokforingsdatum      string `xml:"Folkbokforingsdatum"`
	DistriktKod              int    `xml:"DistriktKod"`
	FolkbokfordForsamlingKod int    `xml:"FolkbokfordForsamlingKod"`
}
type Ns17PersonsokningSvarsPost struct {
	PersonID             PersonID               `xml:"PersonId"`
	Sekretessmarkering   string                 `xml:"Sekretessmarkering"`
	SkyddadFolkbokforing string                 `xml:"SkyddadFolkbokforing"`
	SenasteAndringSPAR   string                 `xml:"SenasteAndringSPAR"`
	Persondetaljer       Persondetaljer         `xml:"Persondetaljer"`
	Folkbokforingsadress []Folkbokforingsadress `xml:"Folkbokforingsadress"`
}
type Ns17SPARPersonsokningSvar struct {
	Ns17PersonsokningSvarsPost Ns17PersonsokningSvarsPost `xml:"ns17:PersonsokningSvarsPost"`
}
type SBody struct {
	Ns17SPARPersonsokningSvar Ns17SPARPersonsokningSvar `xml:"ns17:SPARPersonsokningSvar"`
}
type SEnvelope struct {
	SBody SBody `xml:"S:Body"`
}
