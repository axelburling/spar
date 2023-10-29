package gen

type Person struct {
	PersonID             PersonID
	Sekretessmarkering   string
	SkyddadFolkbokforing string
	SenasteAndringSPAR   string
	Persondetaljer       PersonDetaljer
	Folkbokforingsadress []Folkbokforingsadress
}

type Folkbokforingsadress struct {
	DatumFrom                string
	DatumTill                string
	Utdelningsadress2        string
	PostNr                   int
	Postort                  string
	FolkbokfordLanKod        int
	FolkbokfordKommunKod     int
	Folkbokforingsdatum      string
	DistriktKod              int
	FolkbokfordForsamlingKod int
}

type PersonDetaljer struct {
	DatumFrom            string
	DatumTill            string
	Fornamn              string
	Tilltalsnamn         string
	Efternamn            string
	Sekretessmarkering   string
	SkyddadFolkbokforing string
	Fodelsetid           string
	Kon                  string
}

type PersonID struct {
	FysiskPersonID string
}
