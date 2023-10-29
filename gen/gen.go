package gen

import (
	"math/rand"
	"strconv"
	"time"

	f "github.com/axelburling/spar/formater"
	personnummer "github.com/personnummer/go/v3"
)

func Gen(pin *personnummer.Personnummer) (*f.Ns17PersonsokningSvarsPost, error) {

	pinString, err := pin.Format(true)

	if err != nil {
		return nil, err
	}

	p := f.Ns17PersonsokningSvarsPost{
		PersonID: f.PersonID{
			FysiskPersonID: stringToInt(pinString),
		},
		Sekretessmarkering:   "N",
		SkyddadFolkbokforing: "N",
		SenasteAndringSPAR:   randomPastDate(),
		Persondetaljer:       GenDetails(pin),
		Folkbokforingsadress: GenFolk(3),
	}

	return &p, nil
}

func randomFormList(list []string) string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

func randomPastDate() string {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000

	return time.Unix(randomTime, 0).Format("2006-01-02")
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func GenFolk(l int) []f.Folkbokforingsadress {

	res := []f.Folkbokforingsadress{}

	for i := 0; i < l; i++ {
		fe := f.Folkbokforingsadress{
			DatumFrom:                randomPastDate(),
			DatumTill:                randomPastDate(),
			Utdelningsadress2:        randomFormList([]string{"", "Lgh 123"}),
			PostNr:                   stringToInt(randomFormList(ZipCodes)),
			Postort:                  randomFormList(ZipPlaces),
			DistriktKod:              stringToInt(randomFormList(DistriktCodes)),
			FolkbokfordKommunKod:     stringToInt(randomFormList(StateCodes)),
			FolkbokfordLanKod:        stringToInt(randomFormList(MunicipalityCodes)),
			Folkbokforingsdatum:      randomPastDate(),
			FolkbokfordForsamlingKod: stringToInt(randomFormList(ParishCodes)),
		}

		res = append(res, fe)
	}

	return res
}

func GenDetails(pin *personnummer.Personnummer) f.Persondetaljer {
	g := "K"

	FirstNames := WomenFirstNames

	if pin.IsMale() {
		g = "M"
		FirstNames = MenFirstNames
	}

	return f.Persondetaljer{
		DatumFrom:            randomPastDate(),
		DatumTill:            randomPastDate(),
		Fornamn:              randomFormList(FirstNames),
		Tilltalsnamn:         randomFormList(FirstNames),
		Efternamn:            randomFormList(LastNames),
		Sekretessmarkering:   randomFormList([]string{"N", "J"}),
		SkyddadFolkbokforing: randomFormList([]string{"N", "J"}),
		Fodelsetid:           pin.GetDate().Format("2006-01-02"),
		Kon:                  g,
	}
}
