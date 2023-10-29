package pin

import (
	personnummer "github.com/personnummer/go/v3"
)

func Hello() (*personnummer.Personnummer, error) {
	p, err := personnummer.Parse("1234567890")

	if err != nil {
		return nil, err
	}

	return p, nil
}
